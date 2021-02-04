package main

import (
	"bilibili/function"
	"bilibili/model"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//proxy.XiequTicker()
	for  {
		sibide := 5 * time.Second   //延迟
		function.GetAnio()		//获取源
		var bilibiliAnioModel model.BilibiliAnio
		strArr := [4]string{"我我我","冲啊","来了来了","爱了"}
		rand.Seed(time.Now().UnixNano())
		str := strArr[rand.Intn(len(strArr)-1)]
		data := bilibiliAnioModel.BilibiliAnioList("is_ok = 0")
		for _,v := range data {
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
				function.BilibiliCommentAdd(v.OriginRidStr,strconv.Itoa(types),str)
				if v.PreviousDynamicIdStr != v.OriginDynamicIdStr && v.PreviousDynamicIdStr != ""{
					var types int
					if v.PreviousType == 1 {
						types = 17
					}else if v.PreviousType == 2 {
						types = 11
					} else if v.PreviousType == 8{
						types = 1
					}
					function.BilibiliCommentAdd(v.PreviousRidStr,strconv.Itoa(types),str)
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
		time.Sleep(30 * time.Minute)
	}




}

