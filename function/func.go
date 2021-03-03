package function

import (
	"bilibili/common"
	"bilibili/config"
	"bilibili/logger"
	"bilibili/model"
	//"bilibili/proxy"
	"bilibili/respdata"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	//"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)


func BilibiliLogin() {
	var config = config.GetBilibiliUrl()
	r, err := http.Get(config["LOGIN_CAPTCHA"].(string))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()
	resp, _ := ioutil.ReadAll(r.Body)
	var captcha respdata.BilibiliCaptcha
	json.Unmarshal(resp, &captcha)
	rr, err := http.Get(config["LOGIN_HASH"].(string))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rr.Body.Close()
	respp, _ := ioutil.ReadAll(rr.Body)
	var hash respdata.BilibiliHash
	json.Unmarshal(respp, &hash)

	/*密码拼接*/
	pwd := hash.Hash + config["PWD"].(string)
	rass, err := common.RsaEncrypt([]byte(pwd), []byte(hash.Key))
	if err != nil {
		fmt.Println(err)
		return
	}
	ras := string(rass)
	/*密码拼接*/

	/*登录*/
	c := colly.NewCollector()

	c.OnResponse(func(response *colly.Response) {
		head := response.Headers
		cookie := head.Get("cookie")
		fmt.Println(cookie)
	})

	c.Post(config["LOGIN"].(string), map[string]string{
		"captchaType": "6",
		"username":    config["USER"].(string),
		"password":    ras,
		"keep":        "1",
		"key":         captcha.Data.Result.Key,
		"challenge":   captcha.Data.Result.Challenger,
		"validate":    "",
		"seccode":     "",
	})
	/*登录*/
}

//先获取数据，然后获取详情，关注，转发,评论
//https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?visitor_uid=1268950779&host_uid=18219898&offset_dynamic_id=0&need_top=1&platform=web
func GetAnio() {
	var config = config.GetBilibiliUrl()
	c := colly.NewCollector()

	c.OnResponse(func(response *colly.Response) {
		var list respdata.BilibiliSpaceHistory
		json.Unmarshal(response.Body, &list)
		for _, v := range list.Data.Cards {
			var listCard respdata.BilibiliSpaceHistoryCardJson
			json.Unmarshal([]byte(v.Card), &listCard)
			//uid := strconv.FormatInt(listCard.OriginUser.Info.Uid,10)
			var params model.BilibiliAnio
			params.BilibiliAnioAdd(model.BilibiliAnio{
				OriginDynamicIdStr : v.Desc.OrigDyIdStr,
				OriginRidStr : v.Desc.Origin.RidStr,
				OriginType : v.Desc.Origin.Type,
				OriginUid : v.Desc.Origin.Uid,
				PreviousDynamicIdStr:  v.Desc.Previous.DynamicIdStr,
				PreviousRidStr: v.Desc.Previous.RidStr,
				PreviousType: v.Desc.Previous.Type,
				PreviousUid: v.Desc.Previous.Uid,
				JsonData: common.JsonEncode(v),
				Str: listCard.Item.Content,
				ZhuanfaUid: config["BILIBILI_UID"].(string),
			})
			//fmt.Println("爬取成功,抽奖用户id为："+uid+"名称是")
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(response *colly.Response, err error) {
		logger.LoggerToFile(err.Error())
	})
	c.Visit(config["ANIO"].(string))
}

// 获取互动抽奖列表
func GetBilibiliHuDong()  {

}


// 关注up主
func BilibiliModify(fid []string) {
	var config = config.GetBilibiliUrl()
	c := colly.NewCollector()

	c.OnResponse(func(response *colly.Response) {
		var bilibiliCode respdata.BilibiliCode
		json.Unmarshal(response.Body, &bilibiliCode)
		if bilibiliCode.Code != 0 {
			logger.LoggerToFile("错误：" + bilibiliCode.Message)
			return
		}
		fmt.Println("关注up主成功")
	})
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
		request.Headers.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	})
	for _,v := range fid{
		c.Post(config["UP_MODIFY"].(string), map[string]string{
			"fid":    v,
			"act":    "1",
			"re_src": "11",
			"csrf":   config["CSRF"].(string),
			"jsonp":  "jsonp",
		})
	}

}

//转发动态评论
func BilibiliRepost(dynamicId,str string) {
	var config = config.GetBilibiliUrl()
	c := colly.NewCollector()

	c.OnResponse(func(response *colly.Response) {
		var bilibiliCode respdata.BilibiliCode
		json.Unmarshal(response.Body, &bilibiliCode)
		if bilibiliCode.Code != 0 {
			logger.LoggerToFile("错误：" + bilibiliCode.Message)
			return
		}
		fmt.Println("转发成功")
	})
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
		request.Headers.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	})
	c.Post(config["UP_REPOST"].(string), map[string]string{
		"uid":        config["BILIBILI_UID"].(string),
		"dynamic_id": dynamicId, //动态id
		"content":    str,
		"extension":  "{\"emoji_type\":1}",
		"at_uids":    "",
		"ctrl":       "[]",
		"csrf_token": config["CSRF"].(string),
		"csrf":       config["CSRF"].(string),
	})
}

//图片类型动态评论
func BilibiliCommentAdd(oid,typeStr,str string) {
	var config = config.GetBilibiliUrl()
	c := colly.NewCollector()

	c.OnResponse(func(response *colly.Response) {
		var bilibiliCode respdata.BilibiliCode
		json.Unmarshal(response.Body, &bilibiliCode)
		if bilibiliCode.Code != 0 {
			logger.LoggerToFile("错误：" + bilibiliCode.Message)
			return
		}
		fmt.Println("评论成功")
	})
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
		request.Headers.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	})
	c.Post(config["REPLY"].(string), map[string]string{
		"oid":     oid, //评论区id
		"type":    typeStr,     //评论区类型
		"message": str,
		"plat":    "1",
		"jsonp":   "jsonp",
		"csrf":    config["CSRF"].(string),
	})
}

//
func BilibiliVideoShare(uid,rid,str string) {
	var config = config.GetBilibiliUrl()
	c := colly.NewCollector()

	c.OnResponse(func(response *colly.Response) {
		var bilibiliCode respdata.BilibiliCode
		json.Unmarshal(response.Body, &bilibiliCode)
		if bilibiliCode.Code != 0 {
			logger.LoggerToFile("错误：" + bilibiliCode.Message)
			return
		}
		fmt.Println("分享成功")
	})
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
		request.Headers.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	})
	c.Post(config["VIDEO_REPOST"].(string), map[string]string{
		"rid":     rid,
		"repost_code":     "20000",
		"content":    str,
		"type":    "8",     //评论区类型
		"share_uid": config["BILIBILI_UID"].(string),
		"uid":    uid,
		"platform":   "pc",
		"csrf_token":    config["CSRF"].(string),
	})
}



//获取个人信息
func GetBilibiliUserInfo() respdata.BilibiliUserInfo {
	var config = config.GetBilibiliUrl()
	r, err := http.NewRequest("GET", config["USER_INFO"].(string), nil)
	r.Header.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
	r.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	if err != nil {
		logger.LoggerToFile("错误:" + err.Error())
		return respdata.BilibiliUserInfo{}
	}
	var userInfo respdata.BilibiliUserInfo
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.LoggerToFile("错误:" + err.Error())
		return respdata.BilibiliUserInfo{}
	}
	json.Unmarshal(resp, &userInfo)
	return userInfo
}



//手动抽奖
func BilibiliDoShengYou() {
	var config = config.GetBilibiliUrl()
	c := colly.NewCollector()

	c.OnResponse(func(response *colly.Response) {
		var BilibiliDo respdata.BilibiliDo
		json.Unmarshal(response.Body, &BilibiliDo)
		if BilibiliDo.Code != 0 {
			logger.LoggerToFile("错误：" + BilibiliDo.Message)
			return
		}
		fmt.Println("中奖信息："+BilibiliDo.Data[0].GiftName)
	})
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
		request.Headers.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	})
	c.Post(config["DO_CHOUJIANG"].(string), map[string]string{
		"sid":     "a25363bb-3baf-11eb-8597-246e966235d8",
		"type":     "1",
		"csrf":    config["CSRF"].(string),
	})
}




//获取抽奖列表
func BilibiliGetDo() {
	var config = config.GetBilibiliUrl()
	r,err := http.Get(config["DO_YIFARUHUN"].(string))
	if err != nil {
		logger.LoggerToFile("错误："+err.Error())
		return
	}
	resp,err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var data respdata.BilibiliDynamicDetail
	json.Unmarshal(resp,&data)
	var cardJson = data.Data.Card.CardJson
	json.Unmarshal([]byte(data.Data.Card.Card),&cardJson)
	compile := regexp.MustCompile(`(https?://(www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b[-a-zA-Z0-9()!@:%_+.~#?&//=]*) (.+)`)
	submatch := compile.FindAllStringSubmatch(cardJson.Item.Description, -1)
	for _,v := range submatch {
		if v[3] == "饿了么（测欧）" || v[3] == "三国（积分盘 可攒次数）" || v[3] == "一发入魂盘合集" {
			continue
		}else{
			var BilibliDoAuto model.BilibiliDoAuto
			BilibliDoAuto.BilibiliDoAutoEdit(model.BilibiliDoAuto{
				Url: v[1],
				Name: v[3],
			})
		}
	}
	fmt.Println("获取列表成功")
}


//更新信息
func BilibliDoUpdate() {
	var config = config.GetBilibiliUrl()
	var bilibiliDoAuto  model.BilibiliDoAuto
	data := bilibiliDoAuto.BilibiliDoAutoList()
	client := &http.Client{}
	for _,v := range data {
		var toUid int
		/*获取信息*/
		if v.Sid == "" {
			rr,err := http.Get(v.Url)
			fmt.Println()
			if err != nil {
				logger.LoggerToFile("错误："+err.Error())
				continue
			}
			respp,err := ioutil.ReadAll(rr.Body)
			rr.Body.Close()
			compile := regexp.MustCompile(`window\.__initialState *= *JSON\.parse\(['"]{1}(.+)['"]{1}\);*`)
			submatch := compile.FindAllStringSubmatch(string(respp), -1)
			submatch[0][1] = strings.Replace(submatch[0][1], "\\", "", -1)
			var jsonAct respdata.BilibiliActivity
			json.Unmarshal([]byte(submatch[0][1]),&jsonAct)
			if jsonAct.LotteryNew == nil {
				jsonAct.LotteryNew = jsonAct.PcLotteryNew
			}
			/*获取信息*/
			v.Sid = jsonAct.LotteryNew[0].LotteryId
			v.JsonData = submatch[0][1]

			if jsonAct.FollowNew != nil{
				toUid,err = strconv.Atoi(jsonAct.FollowNew[0].Uid)
			}
			time.Sleep(3*time.Second)
		}



		/*增加关注抽奖机会*/
		var isModify  int
		if v.Sid != "" {
			if v.IsModify == 0 {
				postStr := `sid=`+v.Sid+`&action_type=4&csrf=`+config["CSRF"].(string)
				postReq,_ := http.NewRequest("POST",config["DO_CHOUJIANGNUMADD"].(string),strings.NewReader(postStr))
				postReq.Header.Set("Content-Type","application/x-www-form-urlencoded")
				postReq.Header.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
				postReq.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
				postRep,_ := client.Do(postReq)

				postResp,_ := ioutil.ReadAll(postRep.Body)
				postRep.Body.Close()
				var bilibiliGzAddCode respdata.BilibiliCode
				json.Unmarshal(postResp,&bilibiliGzAddCode)
				if bilibiliGzAddCode.Code != 0 {
					logger.LoggerToFile(v.Sid+"错误提示："+bilibiliGzAddCode.Message)
					fmt.Println(v.Sid+"错误提示："+bilibiliGzAddCode.Message)
				}else{
					fmt.Println("增加关注机会成功")
				}
				isModify = 1
				time.Sleep(3*time.Second)
			}else {
				isModify = 1
			}
			/*增加关注抽奖机会*/

			/*增加分享抽奖机会*/
			//postStrr := `{"sid":"`+v.Sid+`","action_type":3,"csrf":`+config["CSRF"].(string)+`}`
			postStrr := `sid=`+v.Sid+`&action_type=3&csrf=`+config["CSRF"].(string)
			postReqr,_ := http.NewRequest("POST",config["DO_CHOUJIANGNUMADD"].(string),strings.NewReader(postStrr))
			postReqr.Header.Set("Content-Type","application/x-www-form-urlencoded")
			postReqr.Header.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
			postReqr.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
			postRepr,_ := client.Do(postReqr)

			postRespr,_ := ioutil.ReadAll(postRepr.Body)
			postRepr.Body.Close()
			var bilibiliAddCode respdata.BilibiliCode
			json.Unmarshal(postRespr,&bilibiliAddCode)
			if bilibiliAddCode.Code != 0 {
				logger.LoggerToFile(v.Sid+"错误提示："+bilibiliAddCode.Message)
				fmt.Println(v.Sid+"错误提示："+bilibiliAddCode.Message)
			}else{
				fmt.Println("增加分享机会成功")
			}
			time.Sleep(3*time.Second)
			/*增加抽奖机会*/

			/*更新信息*/
			r, _ := http.NewRequest("GET", config["DO_CHOUJIANGNUM"].(string)+v.Sid, nil)
			r.Header.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
			r.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
			rep,_ := client.Do(r)
			resp,_ := ioutil.ReadAll(rep.Body)
			rep.Body.Close()
			var bilibiliChouJiangNumData respdata.BilibiliChouJiangNumData
			json.Unmarshal(resp,&bilibiliChouJiangNumData)
			fmt.Println(bilibiliChouJiangNumData.Data.Times)
			fmt.Println(v.Url)

			v.BilibiliDoAutoEdit(model.BilibiliDoAuto{
				Url: v.Url,
				JsonData: v.JsonData,
				Mid: toUid,
				Sid: v.Sid,
				Num: bilibiliChouJiangNumData.Data.Times,
				IsModify: isModify,
			})
			/*更新信息*/
			fmt.Println("更新数据sid:"+v.Sid+"成功")
		}else{
			fmt.Println("未找到sid")
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println("列表更新完毕")
}


func BilibiliDoRun() {
	var config = config.GetBilibiliUrl()
	var bilibiliMsg model.BilibiliDoMsg

	c := colly.NewCollector()
	c.OnResponse(func(response *colly.Response) {
		var BilibiliDo respdata.BilibiliDo
		json.Unmarshal(response.Body, &BilibiliDo)
		if BilibiliDo.Code != 0 {
			logger.LoggerToFile("错误：" + BilibiliDo.Message)
			return
		}
		bilibiliMsg.Msg = "中奖信息："+BilibiliDo.Data[0].GiftName
		bilibiliMsg.BilibiliDoMsgAdd(bilibiliMsg)
		fmt.Println("中奖信息："+BilibiliDo.Data[0].GiftName)
	})
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("cookie", "SESSDATA="+config["SESSDATA"].(string))
		request.Headers.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	})

	var bilibiliDoAuto  model.BilibiliDoAuto
	data := bilibiliDoAuto.BilibiliDoAutoList()
	for _,v := range data {
		bilibiliMsg.Sid = v.Sid
		bilibiliMsg.Name = v.Name
		if v.Num != 0 {
			for i:=0;i<v.Num;i++{
				c.Post(config["DO_CHOUJIANG"].(string), map[string]string{
					"sid":     v.Sid,
					"type":     "1",
					"csrf":    config["CSRF"].(string),
				})
				fmt.Println(i)
				time.Sleep(6*time.Second)
			}
		}else{
			fmt.Println(v.Name+"没有抽奖机会")
		}
	}
	fmt.Println("所有抽奖完毕")
}