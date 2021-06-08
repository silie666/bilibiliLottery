package main

import (
	"bilibili/config"
	"bilibili/function"
	"bilibili/model"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//proxy.XiequTicker()

	osArr := os.Args

	switch osArr[1] {
	case "-get":
		function.BilibiliGetDo()
		break
	case "-doupdate":
		function.BilibliDoUpdate()
		break
	case "-dorun":
		function.BilibiliDoRun()
		break
	case "-del":
		for true {
			function.BilibiliAutoDel()
			time.Sleep(30*time.Second)
		}
		break
	case "-zhuanfa":
		ticker := time.NewTicker(time.Minute * 5)
		ticker2 := time.NewTicker(time.Minute * 43)
		for {
			select {
			case <-ticker.C:
				function.GetAnio()		//获取源
			case <-ticker2.C:
				sibide := 5 * time.Second   //延迟
				i := 3 //每3条动态
				var config = config.GetBilibiliUrl()
				var bilibiliAnioModel model.BilibiliAnio
				strArr := [4]string{"我我我","不错","来了来了","爱了"}
				ordinaryArr := [3]string{"[doge][doge][doge]","[doge][doge]","[doge]"}
				ordinaryStr := strArr[rand.Intn(len(ordinaryArr)-1)]
				rand.Seed(time.Now().UnixNano())
				str := strArr[rand.Intn(len(strArr)-1)]   //你领到多少红包
				str2 := strArr[rand.Intn(len(strArr)-1)]
				var where = "is_ok = 0 and zhuanfa_uid = " + config["BILIBILI_UID"].(string)
				data := bilibiliAnioModel.BilibiliAnioList(where)
				for _,v := range data {
					if find := strings.Contains(v.Str,"//@"); !find {
						str = v.Str
					}
					var isModify,isRepost,isComment,isLike int
					if v.IsModify == 0 {
						var fid []string
						fid = append(fid,strconv.Itoa(v.OriginUid))
						fid = append(fid,strconv.Itoa(v.PreviousUid))
						function.BilibiliModify(fid)
						isModify = 1
					}else{
						isModify = 1
					}
					time.Sleep(sibide)


					if v.IsRepost == 0 {
						is_false := bilibiliAnioModel.IsBilibiliAnio("origin_dynamic_id_str = "+v.OriginDynamicIdStr+" and zhuanfa_uid = "+v.ZhuanfaUid+" and is_repost = 1")
						if!is_false{
							if v.OriginType == 8 {
								function.BilibiliVideoShare(strconv.Itoa(v.OriginUid),v.OriginRidStr,str)
							}else{
								function.BilibiliRepost(v.OriginDynamicIdStr,str,v.Ctrl)
							}
						}
						if v.PreviousDynamicIdStr != v.OriginDynamicIdStr && v.PreviousDynamicIdStr != ""{
							is_false := bilibiliAnioModel.IsBilibiliAnio("previous_dynamic_id_str = "+v.PreviousDynamicIdStr+" and zhuanfa_uid = "+v.ZhuanfaUid+" and is_repost = 1")
							if !is_false{
								if v.PreviousType == 8 {
									function.BilibiliVideoShare(strconv.Itoa(v.PreviousUid),v.PreviousRidStr,str)
								}else{
									function.BilibiliRepost(v.PreviousDynamicIdStr,v.Str,v.Ctrl)
								}
							}
						}
						isRepost = 1
					}else{
						isRepost = 1
					}
					time.Sleep(sibide)


					if v.IsComment == 0 {
						is_false := bilibiliAnioModel.IsBilibiliAnio("origin_dynamic_id_str = "+v.OriginDynamicIdStr+" and zhuanfa_uid = "+v.ZhuanfaUid+" and is_comment = 1")
						if!is_false{
							var types int
							if v.OriginType == 1 {
								types = 17
							}else if v.OriginType == 2 {
								types = 11
							} else if v.OriginType == 8{
								types = 1
							}
							function.BilibiliCommentAdd(v.OriginRidStr,strconv.Itoa(types),str2)
						}
						if v.PreviousDynamicIdStr != v.OriginDynamicIdStr && v.PreviousDynamicIdStr != ""{
							is_false := bilibiliAnioModel.IsBilibiliAnio("previous_dynamic_id_str = "+v.PreviousDynamicIdStr+" and zhuanfa_uid = "+v.ZhuanfaUid+" and is_comment = 1")
							fmt.Println(is_false)
							if !is_false{
								var types int
								if v.PreviousType == 1 {
									types = 17
								}else if v.PreviousType == 2 {
									types = 11
								} else if v.PreviousType == 8{
									types = 1
								}
								fmt.Println(v.PreviousDynamicIdStr,strconv.Itoa(types),str2)
								function.BilibiliCommentAdd(v.PreviousDynamicIdStr,strconv.Itoa(types),str2)
							}
						}
						isComment = 1
					}else{
						isComment = 1
					}

					time.Sleep(sibide)
					if v.IsLike == 0 {
						is_false := bilibiliAnioModel.IsBilibiliAnio("origin_dynamic_id_str = "+v.OriginDynamicIdStr+" and zhuanfa_uid = "+v.ZhuanfaUid+" and is_like = 1")
						if !is_false{
							if v.OriginType == 1 || v.OriginType == 2{
								function.BilibiliDynamicLike(v.OriginDynamicIdStr)
							}else if v.OriginType == 8{
								function.BilibiliVideoLike(v.Bvid)
							}
						}
						if v.PreviousDynamicIdStr != v.OriginDynamicIdStr && v.PreviousDynamicIdStr != ""{
							is_false := bilibiliAnioModel.IsBilibiliAnio("previous_dynamic_id_str = "+v.PreviousDynamicIdStr+" and zhuanfa_uid = "+v.ZhuanfaUid+" and is_like = 1")
							if !is_false{
								if v.OriginType == 1 || v.OriginType == 2{
									function.BilibiliDynamicLike(v.PreviousDynamicIdStr)
								}else if v.OriginType == 8{
									function.BilibiliVideoLike(v.Bvid)
								}
							}
						}
						isLike = 1
					} else {
						isLike = 1
					}
					bilibiliAnioModel.BilibiliAnioSave(model.BilibiliAnio{
						Id: v.Id,
						IsComment: isComment,
						IsModify: isModify,
						IsRepost: isRepost,
						IsLike: isLike,
						IsOk: 1,
					})
					time.Sleep(sibide)
					i++
					if i % 3 == 0 {  //每3条动态发一条普通动态，防止过滤
						function.BilibiliOrdinary(ordinaryStr)
					}
				}
			}
		}
	default:
		fmt.Println("请输入命令，-get获取抽奖列表，-dorun开始抽奖，-del删除动态，-zhuanfa开始转发动态抽奖")
	}


}


