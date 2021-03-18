package main

import (
	"bilibili/config"
	"bilibili/function"
	"bilibili/model"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	//proxy.XiequTicker()

	ticker := time.NewTicker(time.Minute * 5)
	ticker2 := time.NewTicker(time.Minute * 43)
	for {
		select {
		case <-ticker.C:
			function.GetAnio()		//获取源
		case <-ticker2.C:
			sibide := 5 * time.Second   //延迟
			var config = config.GetBilibiliUrl()
			var bilibiliAnioModel model.BilibiliAnio
			strArr := [4]string{"我我我","不错","来了来了","爱了"}
			rand.Seed(time.Now().UnixNano())
			str := strArr[rand.Intn(len(strArr)-1)]   //你领到多少红包
			str2 := strArr[rand.Intn(len(strArr)-1)]
			var where = "is_ok = 0 and zhuanfa_uid = " + config["BILIBILI_UID"].(string)
			data := bilibiliAnioModel.BilibiliAnioList(where)
			for _,v := range data {
				if find := strings.Contains(v.Str,"//@"); !find {
					str = v.Str
				}
				var isModify,isRepost,isComment int
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
					if v.OriginType == 8 {
						function.BilibiliVideoShare(strconv.Itoa(v.OriginUid),v.OriginRidStr,str)
					}else{
						function.BilibiliRepost(v.OriginDynamicIdStr,str)
					}
					if v.PreviousDynamicIdStr != v.OriginDynamicIdStr && v.PreviousDynamicIdStr != ""{
						if v.PreviousType == 8 {
							function.BilibiliVideoShare(strconv.Itoa(v.PreviousUid),v.PreviousRidStr,str)
						}else{
							function.BilibiliRepost(v.PreviousDynamicIdStr,v.Str)
						}
					}
					isRepost = 1
				}else{
					isRepost = 1
				}
				time.Sleep(sibide)


				if v.IsComment == 0 {
					var types int
					if v.OriginType == 1 {
						types = 17
					}else if v.OriginType == 2 {
						types = 11
					} else if v.OriginType == 8{
						types = 1
					}
					function.BilibiliCommentAdd(v.OriginRidStr,strconv.Itoa(types),str2)
					if v.PreviousDynamicIdStr != v.OriginDynamicIdStr && v.PreviousDynamicIdStr != ""{
						var types int
						if v.PreviousType == 1 {
							types = 17
						}else if v.PreviousType == 2 {
							types = 11
						} else if v.PreviousType == 8{
							types = 1
						}
						function.BilibiliCommentAdd(v.PreviousRidStr,strconv.Itoa(types),str2)
					}
					isComment = 1
				}else{
					isComment = 1
				}
				bilibiliAnioModel.BilibiliAnioSave(model.BilibiliAnio{
					Id: v.Id,
					IsComment: isComment,
					IsModify: isModify,
					IsRepost: isRepost,
					IsOk: 1,
				})
				time.Sleep(sibide)
			}
		}
	}
}


