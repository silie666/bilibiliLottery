package config


func GetBilibiliUrl() map[string]interface{} {
	// 初始化数据库配置map
	urlConfig := make(map[string]interface{})
	urlConfig["USER"] = "*****"
	urlConfig["PWD"] = "*****"

	//我
	urlConfig["SESSDATA"] = "6ce09859%2C1627113791%2C0d0df*11"
	urlConfig["CSRF"] = "3c9344bffeb3380d71b2aae73036c6d5"
	urlConfig["BILIBILI_UID"] = "1268950779"

	//蛇
	//urlConfig["SESSDATA"] = "8b896bfb%2C1630129369%2C4165e%2A31"
	//urlConfig["CSRF"] = "b99683140cd0a77b1c20de2e77eeab09"
	//urlConfig["BILIBILI_UID"] = "400141956"

	//雕
	//urlConfig["SESSDATA"] = "93f1d9dd%2C1630130133%2Cf7f19%2A31"
	//urlConfig["CSRF"] = "3401913346d0d87ce24cfecab14e9b2f"
	//urlConfig["BILIBILI_UID"] = "1662242662"

	//企鹅
	//urlConfig["SESSDATA"] = "fcea3399%2C1630132168%2C28af8%2A31 "
	//urlConfig["CSRF"] = "85c91764dabdd544a5c652013d2dac52"
	//urlConfig["BILIBILI_UID"] = "48137033"

	//企鹅2
	//urlConfig["SESSDATA"] = "79029506%2C1630209525%2C06339%2A31"
	//urlConfig["CSRF"] = "2939daa0f71d0fa5fd9fb30a91bc767e"
	//urlConfig["BILIBILI_UID"] = "232971386"

	//企鹅3
	//urlConfig["SESSDATA"] = "700c0d34%2C1630145238%2Cc655b*31"
	//urlConfig["CSRF"] = "28f4354491f695f70fa604145de0996c"
	//urlConfig["BILIBILI_UID"] = "523455358"






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



	urlConfig["DO_YIFARUHUN"] = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/get_dynamic_detail?dynamic_id=494287213211523960"  //485011372337648035最新 483153463975687658一发入魂动态id
	urlConfig["DO_CHOUJIANG"] = "https://api.bilibili.com/x/activity/lottery/do"  //抽奖链接
	urlConfig["DO_CHOUJIANGNUM"] = "https://api.bilibili.com/x/activity/lottery/mytimes?sid="  //查询抽奖次数
	urlConfig["DO_CHOUJIANGNUMADD"] = "https://api.bilibili.com/x/activity/lottery/addtimes"  //增加抽奖次数



	//拜年活动
	urlConfig["DO_BNJ"] = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/get_dynamic_detail?dynamic_id=489463763067784879"  //485011372337648035最新 483153463975687658一发入魂动态id





	return urlConfig
}
