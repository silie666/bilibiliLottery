package config


func GetBilibiliUrl() map[string]interface{} {
	// 初始化数据库配置map
	urlConfig := make(map[string]interface{})
	urlConfig["USER"] = "*****"
	urlConfig["PWD"] = "*****"


	urlConfig["SESSDATA"] = "6ce09859%2C1627113791%2C0d0df*11"
	urlConfig["CSRF"] = "3c9344bffeb3380d71b2aae73036c6d5"
	urlConfig["BILIBILI_UID"] = "1268950779"






	urlConfig["LOGIN"] = "http://passport.bilibili.com/web/login/v2"
	urlConfig["LOGIN_CAPTCHA"] = "https://passport.bilibili.com/web/captcha/combine?plat=6"
	urlConfig["LOGIN_HASH"] = "http://passport.bilibili.com/login?act=getkey"


	urlConfig["ANIO"] = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?visitor_uid=1268950779&host_uid=18219898&offset_dynamic_id=0&need_top=1&platform=web"
	urlConfig["UP_MODIFY"] = "https://api.bilibili.com/x/relation/modify"
	urlConfig["UP_REPOST"] = "https://api.vc.bilibili.com/dynamic_repost/v1/dynamic_repost/repost"
	urlConfig["USER_INFO"] = "http://api.bilibili.com/x/web-interface/nav"
	urlConfig["REPLY"] = "https://api.bilibili.com/x/v2/reply/add"

	urlConfig["VIDEO_REPOST"] = "https://api.vc.bilibili.com/dynamic_repost/v1/dynamic_repost/share"


	//urlConfig["HUDONG_HOT"] = "https://api.vc.bilibili.com/topic_svr/v1/topic_svr/topic_new?topic_id=3230836"
	//urlConfig["HUDONG_NEW"] = "https://api.vc.bilibili.com/topic_svr/v1/topic_svr/topic_history?topic_name=%E4%BA%92%E5%8A%A8%E6%8A%BD%E5%A5%96&offset_dynamic_id=" //484787016136660229



	urlConfig["DO_YIFARUHUN"] = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/get_dynamic_detail?dynamic_id=483153463975687658"  //485011372337648035最新 483153463975687658一发入魂动态id
	urlConfig["DO_CHOUJIANG"] = "https://api.bilibili.com/x/activity/lottery/do"  //抽奖链接
	urlConfig["DO_CHOUJIANGNUM"] = "https://api.bilibili.com/x/activity/lottery/mytimes?sid="  //查询抽奖次数
	urlConfig["DO_CHOUJIANGNUMADD"] = "https://api.bilibili.com/x/activity/lottery/addtimes"  //增加抽奖次数




	return urlConfig
}
