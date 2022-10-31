package function

import (
	"bilibili/common"
	"bilibili/config"
	"bilibili/drivers/mysql"
	"bilibili/logger"
	"bilibili/model"
	"bilibili/respdata"
	"encoding/json"
	"fmt"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var Env = config.Env

func Init() {
	if !mysql.Db.Migrator().HasTable(&model.BilibiliAnio{}) {
		mysql.Db.Migrator().CreateTable(&model.BilibiliAnio{})
	}
	if !mysql.Db.Migrator().HasTable(&model.BilibiliDoAuto{}) {
		mysql.Db.Migrator().CreateTable(&model.BilibiliDoAuto{})
	}
	if !mysql.Db.Migrator().HasTable(&model.BilibiliDoMsg{}) {
		mysql.Db.Migrator().CreateTable(&model.BilibiliDoMsg{})
	}
	if mysql.Db.Migrator().HasColumn(&model.BilibiliAnio{}, "zhuanfa_uid") {
		mysql.Db.Migrator().RenameColumn(&model.BilibiliAnio{}, "zhuanfa_uid", "forward_uid")
	}
	fmt.Println("初始化成功")
}

func BilibiliLogin() {
	var userInfo respdata.BilibiliUserInfo
	if Env.GetString("cookie.sess_data") != "" {
		userInfo = GetBilibiliUserInfo()
	}
	if userInfo.Mid == "" {
		fmt.Println("回车显示登录二维码")
		fmt.Scanf("%s", "")
		var bilibiliQRCodeData respdata.BilibiliQRCode
		bilibiliQRCodeData = GetQRCode()
		qrcode := qrcodeTerminal.New()
		qrcode.Get([]byte(bilibiliQRCodeData.Data.Url)).Print()
		fmt.Println("或将此链接复制到手机B站打开:", bilibiliQRCodeData.Data.Url)
		VerifyLogin(bilibiliQRCodeData.Data.OauthKey)
	}
}

func GetQRCode() (qrcode respdata.BilibiliQRCode) {
	response := common.Get(Env.GetString("api.login"), nil)
	err := json.Unmarshal([]byte(response), &qrcode)
	if err != nil {
		panic(err)
	}
	if qrcode.Code != 0 {
		panic("接口错误:" + response)
	}
	return
}

func VerifyLogin(oauthKey string) {
	isLogin := true
	for isLogin {
		time.Sleep(time.Second * 3)

		params := url.Values{}
		params.Add("oauthKey", oauthKey)
		response := common.Post(Env.GetString("api.login_info"), params, "application/x-www-form-urlencoded")
		jsonDecode := common.JsonDecode([]byte(response))
		status := jsonDecode.Get("status").MustBool()
		if status != false {
			responseUrl, _ := jsonDecode.Get("data").Get("url").String()
			loginUrl, _ := url.Parse(responseUrl)
			Env.Set("cookie.sess_data", loginUrl.Query().Get("SESSDATA"))
			Env.Set("cookie.csrf", loginUrl.Query().Get("bili_jct"))
			Env.Set("cookie.uid", loginUrl.Query().Get("DedeUserID"))
			err := Env.WriteConfig()
			if err != nil {
				panic(err)
			}
			fmt.Println("登录成功，运行转发或者抽奖开始")
			isLogin = false
		} else {
			fmt.Println(response)
		}
	}
}

// 同步动态
func SyncDynamic() {
	for _, v := range strings.Split(Env.GetString("data.host_uid"), ",") {
		fmt.Println(v)
		GetDynamic(v)
	}
}

// 获取动态
func GetDynamic(hostUid string) {
	query := url.Values{}
	query.Add("host_uid", hostUid)
	query.Add("offset_dynamic_id", "0")
	query.Add("need_top", "1")
	query.Add("platform", "web")
	response := common.Get(Env.GetString("api.space"), query)
	var list respdata.BilibiliSpaceHistory
	json.Unmarshal([]byte(response), &list)
	for _, v := range list.Data.Cards {
		var listCard respdata.BilibiliSpaceHistoryCardJson
		json.Unmarshal([]byte(v.Card), &listCard)
		var extend respdata.BilibiliExtendJson
		json.Unmarshal([]byte(v.ExtendJson), &extend)
		ctrl, _ := json.Marshal(extend.Ctrl)
		var anio model.BilibiliAnio
		anio.BilibiliAnioAdd(model.BilibiliAnio{
			OriginDynamicIdStr:   v.Desc.OrigDyIdStr,
			OriginRidStr:         v.Desc.Origin.RidStr,
			OriginType:           v.Desc.Origin.Type,
			OriginUid:            v.Desc.Origin.Uid,
			Bvid:                 v.Desc.Origin.Bvid,
			PreviousDynamicIdStr: v.Desc.Previous.DynamicIdStr,
			PreviousRidStr:       v.Desc.Previous.RidStr,
			PreviousType:         v.Desc.Previous.Type,
			PreviousUid:          v.Desc.Previous.Uid,
			JsonData:             common.JsonEncode(v),
			Str:                  listCard.Item.Content,
			ForwardUid:           Env.GetString("cookie.uid"),
			Ctrl:                 string(ctrl),
		})
	}
}

// 关注或取关
func Modify(fid, act string) {
	params := url.Values{}
	params.Add("fid", fid)
	params.Add("act", act)
	params.Add("re_src", "11")
	params.Add("csrf", Env.GetString("cookie.csrf"))
	params.Add("jsonp", "jsonp")
	response := common.Post(Env.GetString("api.modify"), params, "application/x-www-form-urlencoded")
	err := common.BilibiliIsError(response)
	if err != nil {
		return
	}
	if act == "1" {
		fmt.Println("关注成功")
	} else if act == "2" {
		fmt.Println("取关成功")
	}
}

// 批量关注up主
func BilibiliModify(fid []string) {
	for _, v := range fid {
		Modify(v, "1")
	}
}

// 转发
func BilibiliReply(uid, rid, content string) {
	params := url.Values{}
	params.Add("uid", uid)
	params.Add("type", "1")
	params.Add("rid", rid)
	params.Add("content", content)
	params.Add("repost_code", "30000")
	params.Add("from", "create.comment")
	params.Add("extension", "{\"emoji_type\":1}")
	response := common.Post(Env.GetString("api.reply"), params, "application/x-www-form-urlencoded")
	err := common.BilibiliIsError(response)
	if err != nil {
		return
	}
	fmt.Println("转发成功")
}

// 动态点赞
func BilibiliDynamicLike(dynamicId string) {
	params := url.Values{}
	params.Add("uid", Env.GetString("uid"))
	params.Add("dynamic_id", dynamicId)
	params.Add("up", "1")
	params.Add("csrf", Env.GetString("cookie.csrf"))
	response := common.Post(Env.GetString("api.like_dynamic"), params, "application/x-www-form-urlencoded")
	err := common.BilibiliIsError(response)
	if err != nil {
		return
	}
	fmt.Println("点赞成功")
}

// 视频点赞
func BilibiliVideoLike(bvid string) {
	params := url.Values{}
	params.Add("bvid", bvid)
	params.Add("like", "1")
	params.Add("csrf", Env.GetString("cookie.csrf"))
	response := common.Post(Env.GetString("api.like_video"), params, "application/x-www-form-urlencoded")
	err := common.BilibiliIsError(response)
	if err != nil {
		return
	}
	fmt.Println("点赞成功")
}

// 评论
func BilibiliCommentAdd(oid, typeStr, str string) {
	params := url.Values{}
	params.Add("oid", oid)
	params.Add("type", typeStr)
	params.Add("message", str)
	params.Add("plat", "1")
	params.Add("ordering", "heat")
	params.Add("jsonp", "jsonp")
	params.Add("csrf", Env.GetString("cookie.csrf"))
	response := common.Post(Env.GetString("api.comment"), params, "application/x-www-form-urlencoded")
	err := common.BilibiliIsError(response)
	if err != nil {
		return
	}
	fmt.Println("评论成功")
}

// 发布动态
func BilibiliDyn(contents []respdata.DynContent, rid string, scene int) {
	params := respdata.DynBody{}
	params.DynReq.Content.Contents = contents
	params.DynReq.Scene = scene
	params.WebRepostSrc.DynIdStr = rid
	response := common.Post(Env.GetString("api.dyn")+Env.GetString("cookie.csrf"), params, "application/json;charset=UTF-8")
	err := common.BilibiliIsError(response)
	if err != nil {
		return
	}
	fmt.Println("发布成功")
}

// 获取个人信息
func GetBilibiliUserInfo() (userInfo respdata.BilibiliUserInfo) {
	response := common.Get(Env.GetString("api.info"), nil)
	err := common.BilibiliIsError(response)
	if err != nil {
		Env.Set("cookie.sess_data", "")
		Env.Set("cookie.csrf", "")
		Env.Set("cookie.uid", "")
		Env.WriteConfig()
		return respdata.BilibiliUserInfo{}
	}
	json.Unmarshal([]byte(response), &userInfo)
	return
}

// 抽奖
func BilibiliDo(sid string) {
	query := url.Values{}
	query.Add("sid", sid)
	query.Add("type", "1")
	query.Add("csrf", Env.GetString("cookie.csrf"))
	response := common.Get(Env.GetString("api.do_luck_draw"), query)
	err := common.BilibiliIsError(response)
	if err != nil {
		return
	}
}

func BilibiliGeTLuckDraw() {
	mysql.Db.Where("1=1").Delete(&model.BilibiliDoAuto{})
	query := url.Values{}
	query.Add("host_uid", Env.GetString("data.luck_draw_uid"))
	query.Add("offset_dynamic_id", "0")
	query.Add("need_top", "1")
	query.Add("platform", "web")
	response := common.Get(Env.GetString("api.space"), query)
	var list respdata.BilibiliSpaceHistory
	json.Unmarshal([]byte(response), &list)
	for _, v := range list.Data.Cards {
		if v.Display.TopicInfo.NewTopic.Id == Env.GetInt("data.luck_draw_topic_id") {
			var extend respdata.BilibiliExtendJson
			extendJson := v.ExtendJson[:2] + "data" + v.ExtendJson[2:]
			json.Unmarshal([]byte(extendJson), &extend)
			compile := regexp.MustCompile(`(https?://(www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b[-a-zA-Z0-9()!@:%_+.~#?&//=]*) (\S{0,10})`)
			submatch := compile.FindAllStringSubmatch(extend.Data.Content, -1)
			for _, vv := range submatch {
				var BilibliDoAuto model.BilibiliDoAuto
				BilibliDoAuto.BilibiliDoAutoEdit(model.BilibiliDoAuto{
					Url:  vv[1],
					Name: vv[3],
				})
			}
			fmt.Println("获取列表成功")
			return
		}
	}
	fmt.Println("获取列表失败")
}

// 更新信息
func BilibliDoUpdate() {
	var bilibiliDoAuto model.BilibiliDoAuto
	data := bilibiliDoAuto.BilibiliDoAutoList()
	for _, v := range data {
		var toUid int
		/*获取信息*/
		if v.Sid == "" {
			shortUrlRedirect := common.ShortUrlRedirect(v.Url)
			response := common.Get(shortUrlRedirect, nil)
			compile := regexp.MustCompile(`window\.__initialState *= ({.+});`)
			submatch := compile.FindAllStringSubmatch(response, -1)
			if submatch == nil {
				bilibiliDoAuto.BilibiliDoDel(v.Id)
				logger.LogToFile("信息获取失败，活动网页为:" + v.Url)
				continue
			}
			actChar := strings.Index(submatch[0][1], "BaseInfo")
			if actChar == -1 {
				bilibiliDoAuto.BilibiliDoDel(v.Id)
				logger.LogToFile("信息获取失败，活动网页为:" + v.Url)
				continue
			}
			var jsonAct respdata.BilibiliActivity
			json.Unmarshal([]byte(submatch[0][1]), &jsonAct)
			if jsonAct.LotteryNew == nil {
				jsonAct.LotteryNew = jsonAct.PcLotteryNew
			}
			if jsonAct.LotteryNew == nil {
				jsonAct.LotteryNew = jsonAct.H5LotteryV3
			}
			if jsonAct.LotteryNew == nil {
				jsonAct.LotteryNew = jsonAct.PcLotteryV3
			}
			/*获取信息*/
			v.Sid = jsonAct.LotteryNew[0].LotteryId
			v.JsonData = submatch[0][1]

			if jsonAct.FollowNew != nil {
				toUid, _ = strconv.Atoi(jsonAct.FollowNew[0].Uid)
			}
			time.Sleep(3 * time.Second)
		}

		/*增加关注抽奖机会*/
		var isModify int
		if v.Sid != "" {
			var response string
			params := url.Values{}
			params.Add("sid", v.Sid)
			params.Add("action_type", "4")
			params.Add("csrf", Env.GetString("csrf"))
			if v.IsModify == 0 {
				response = common.Post(Env.GetString("api.do_luck_draw_num_add"), params, "application/x-www-form-urlencoded")
				var bilibiliGzAddCode respdata.BilibiliCode
				json.Unmarshal([]byte(response), &bilibiliGzAddCode)
				if bilibiliGzAddCode.Code != 0 {
					logger.LogToFile(v.Sid + "错误提示：" + bilibiliGzAddCode.Message)
				} else {
					fmt.Println("增加关注机会成功")
				}
				isModify = 1
				time.Sleep(3 * time.Second)
			} else {
				isModify = 1
			}
			/*增加关注抽奖机会*/

			/*增加分享抽奖机会*/
			params.Set("action_type", "3")
			response = common.Post(Env.GetString("api.do_luck_draw_num_add"), params, "application/x-www-form-urlencoded")
			var bilibiliAddCode respdata.BilibiliCode
			json.Unmarshal([]byte(response), &bilibiliAddCode)
			if bilibiliAddCode.Code != 0 {
				logger.LogToFile(v.Sid + "错误提示：" + bilibiliAddCode.Message)
			} else {
				fmt.Println("增加分享机会成功")
			}
			time.Sleep(3 * time.Second)
			/*增加抽奖机会*/

			/*更新信息*/
			params.Del("action_type")
			response = common.Get(Env.GetString("api.do_luck_draw_num"), params)
			var bilibiliChouJiangNumData respdata.BilibiliChouJiangNumData
			json.Unmarshal([]byte(response), &bilibiliChouJiangNumData)
			fmt.Println(bilibiliChouJiangNumData.Data.Times)
			fmt.Println(v.Url)

			v.BilibiliDoAutoEdit(model.BilibiliDoAuto{
				Url:      v.Url,
				JsonData: v.JsonData,
				Mid:      toUid,
				Sid:      v.Sid,
				Num:      bilibiliChouJiangNumData.Data.Times,
				IsModify: isModify,
				Name:     v.Name,
			})
			/*更新信息*/
			fmt.Println("更新数据sid:" + v.Sid + "成功")
		} else {
			fmt.Println("未找到sid")
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println("列表更新完毕")
}

func BilibiliDoRun() {
	var bilibiliMsg model.BilibiliDoMsg
	var bilibiliDoAuto model.BilibiliDoAuto
	data := bilibiliDoAuto.BilibiliDoAutoList()
	for _, v := range data {
		bilibiliMsg.Sid = v.Sid
		bilibiliMsg.Name = v.Name
		if v.Num != 0 {
			for i := 0; i < v.Num; i++ {
				params := url.Values{}
				params.Add("sid", v.Sid)
				params.Add("type", "1")
				params.Add("csrf", Env.GetString("csrf"))
				response := common.Post(Env.GetString("api.do_luck_draw"), params, "application/x-www-form-urlencoded")
				var BilibiliDo respdata.BilibiliDo
				json.Unmarshal([]byte(response), &BilibiliDo)
				if BilibiliDo.Code != 0 {
					logger.LogToFile("错误：" + BilibiliDo.Message)
					return
				}
				bilibiliMsg.Msg = "中奖信息：" + BilibiliDo.Data[0].GiftName
				bilibiliMsg.BilibiliDoMsgAdd(bilibiliMsg)
				fmt.Println("中奖信息：" + BilibiliDo.Data[0].GiftName)
				time.Sleep(6 * time.Second)
			}
		} else {
			fmt.Println(v.Name + "没有抽奖机会")
		}
	}
	fmt.Println("所有抽奖完毕")
}

// 删除第二页第一条动态
func BilibiliAutoDel() {
	query := url.Values{}
	query.Add("host_uid", Env.GetString("cookie.uid"))
	query.Add("offset_dynamic_id", "0")
	query.Add("need_top", "1")
	query.Add("platform", "web")
	response := common.Get(Env.GetString("api.space"), query)
	var list respdata.BilibiliSpaceHistory
	json.Unmarshal([]byte(response), &list)
	dynamic_id := strconv.Itoa(list.Data.NextOffset)

	params := url.Values{}
	params.Add("dynamic_id", dynamic_id)
	params.Add("csrf", Env.GetString("csrf"))
	response = common.Post(Env.GetString("api.del_dynamic"), params, "application/x-www-form-urlencoded")
	err := common.BilibiliIsError(response)
	if err != nil {
		return
	}
	fmt.Println("删除成功:", dynamic_id)
}

// 防止过滤，发送一条普通动态
func BilibiliOrdinary() {
	ordinaryArr := [3]string{"[doge][doge][doge]", "[doge][doge]", "[doge]"}
	ordinaryStr := ordinaryArr[rand.Intn(len(ordinaryArr)-1)]
	var ordinaryContents = make([]respdata.DynContent, 1, 10)
	ordinaryContents[0].RawText = ordinaryStr
	ordinaryContents[0].BizId = ""
	ordinaryContents[0].Type = 9
	BilibiliDyn(ordinaryContents, "", 1)
}

func BilibiliStart() {
	ticker := time.NewTicker(time.Minute * Env.GetDuration("data.sync_dynamic_interval"))
	ticker2 := time.NewTicker(time.Minute * Env.GetDuration("data.forward_interval"))
	for {
		select {
		case <-ticker.C:
			SyncDynamic()
		case <-ticker2.C:
			BilibiliForward()
		}
	}
}

// 开始转发
func BilibiliForward() {

	delay := 5 * time.Second //延迟
	i := 3                   //每3条动态
	var bilibiliAnioModel model.BilibiliAnio
	strArr := [4]string{"我我我", "不错", "来了来了", "爱了"}

	ordinaryArr := [3]string{"[doge][doge][doge]", "[doge][doge]", "[doge]"}
	ordinaryStr := ordinaryArr[rand.Intn(len(ordinaryArr)-1)]
	var ordinaryContents = make([]respdata.DynContent, 1, 10)
	ordinaryContents[0].RawText = ordinaryStr
	ordinaryContents[0].BizId = ""
	ordinaryContents[0].Type = 9

	rand.Seed(time.Now().UnixNano())
	str := strArr[rand.Intn(len(strArr)-1)]
	var where = "is_ok = 0 and forward_uid = " + Env.GetString("cookie.uid")
	data := bilibiliAnioModel.BilibiliAnioList(where)
	for _, v := range data {
		var isModify, isRepost, isLike, isComment int

		if v.IsModify == 0 {
			var fid []string
			fid = append(fid, strconv.Itoa(v.OriginUid))
			if v.PreviousUid != 0 {
				fid = append(fid, strconv.Itoa(v.PreviousUid))
			}
			if v.Ctrl != "" {
				var uids []respdata.Ctrl
				json.Unmarshal([]byte(v.Ctrl), &uids)
				for _, vv := range uids {
					if vv.Data == strconv.Itoa(v.OriginUid) || vv.Data == strconv.Itoa(v.PreviousUid) {
						continue
					}
					fid = append(fid, vv.Data)
				}
			}
			BilibiliModify(fid)
			isModify = 1
		} else {
			isModify = 1
		}
		time.Sleep(delay)

		if v.IsRepost == 0 {
			is_false := bilibiliAnioModel.IsBilibiliAnio("origin_dynamic_id_str = " + v.OriginDynamicIdStr + " and forward_uid = " + v.ForwardUid + " and is_repost = 1")
			if !is_false {
				BilibiliReply(strconv.Itoa(v.OriginUid), v.OriginDynamicIdStr, str)
			}
			if v.PreviousDynamicIdStr != v.OriginDynamicIdStr && v.PreviousDynamicIdStr != "" {
				is_false = bilibiliAnioModel.IsBilibiliAnio("previous_dynamic_id_str = " + v.PreviousDynamicIdStr + " and forward_uid = " + v.ForwardUid + " and is_repost = 1")
				if !is_false {
					BilibiliReply(strconv.Itoa(v.PreviousUid), v.PreviousDynamicIdStr, str)
				}
			}
			isRepost = 1
		} else {
			isRepost = 1
		}
		time.Sleep(delay)

		if v.IsComment == 0 {
			is_false := bilibiliAnioModel.IsBilibiliAnio("origin_dynamic_id_str = " + v.OriginDynamicIdStr + " and forward_uid = " + v.ForwardUid + " and is_comment = 1")
			if !is_false {
				var types int
				if v.OriginType == 1 {
					types = 17
				} else if v.OriginType == 2 {
					types = 11
				} else if v.OriginType == 8 {
					types = 1
				}
				BilibiliCommentAdd(v.OriginRidStr, strconv.Itoa(types), str)
			}
			if v.PreviousDynamicIdStr != v.OriginDynamicIdStr && v.PreviousDynamicIdStr != "" {
				is_false := bilibiliAnioModel.IsBilibiliAnio("previous_dynamic_id_str = " + v.PreviousDynamicIdStr + " and forward_uid = " + v.ForwardUid + " and is_comment = 1")
				if !is_false {
					var types int
					if v.PreviousType == 1 {
						types = 17
					} else if v.PreviousType == 2 {
						types = 11
					} else if v.PreviousType == 8 {
						types = 1
					}
					BilibiliCommentAdd(v.PreviousDynamicIdStr, strconv.Itoa(types), str)
				}
			}
			isComment = 1
		} else {
			isComment = 1
		}

		if v.IsLike == 0 {
			is_false := bilibiliAnioModel.IsBilibiliAnio("origin_dynamic_id_str = " + v.OriginDynamicIdStr + " and forward_uid = " + v.ForwardUid + " and is_like = 1")
			if !is_false {
				if v.OriginType == 1 || v.OriginType == 2 {
					BilibiliDynamicLike(v.OriginDynamicIdStr)
				} else if v.OriginType == 8 {
					BilibiliVideoLike(v.Bvid)
				}
			}
			if v.PreviousDynamicIdStr != v.OriginDynamicIdStr && v.PreviousDynamicIdStr != "" {
				is_false := bilibiliAnioModel.IsBilibiliAnio("previous_dynamic_id_str = " + v.PreviousDynamicIdStr + " and forward_uid = " + v.ForwardUid + " and is_like = 1")
				if !is_false {
					if v.OriginType == 1 || v.OriginType == 2 {
						BilibiliDynamicLike(v.PreviousDynamicIdStr)
					} else if v.OriginType == 8 {
						BilibiliVideoLike(v.Bvid)
					}
				}
			}
			isLike = 1
		} else {
			isLike = 1
		}
		bilibiliAnioModel.BilibiliAnioSave(model.BilibiliAnio{
			Id:        v.Id,
			IsModify:  isModify,
			IsRepost:  isRepost,
			IsComment: isComment,
			IsLike:    isLike,
			IsOk:      1,
		})
		time.Sleep(delay)
		i++
		if i%3 == 0 { //每3条动态发一条普通动态，防止过滤
			BilibiliDyn(ordinaryContents, "", 1)
		}
	}
}

// 取消关注
func BilibiliCancelModify() {
	query := url.Values{}
	query.Add("vmid", Env.GetString("cookie.uid"))
	query.Add("order", "asc")
	response := common.Get(Env.GetString("api.modify_list"), query)
	var list respdata.ModifyList
	json.Unmarshal([]byte(response), &list)
	for _, v := range list.Data.List {
		Modify(strconv.Itoa(v.Mid), "2")
		time.Sleep(3 * time.Second)
	}
	BilibiliCancelModify()
}
