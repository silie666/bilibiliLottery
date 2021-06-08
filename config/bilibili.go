package config


func GetBilibiliUrl() map[string]interface{} {
	// 初始化数据库配置map
	urlConfig := make(map[string]interface{})
	urlConfig["USER"] = "*****"  //无视
	urlConfig["PWD"] = "*****"	//无视

	urlConfig["LOGS_PATH"] = "E:/20210425/4f466b1a-3b05-401a-b3de-1c0fa0651f23/wyf/go/src/bilibiliLottery/logger/logs.log"	//日志路径
	//我
	urlConfig["SESSDATA"] = "7437f6ba%2C1635305552%2C91435%2A41"   				//b站cookei SESSDATA
	urlConfig["BUVID3"] = "55AC3C2A-085D-4C00-A3A5-D6934B3D6E9E13414infoc"		//b站cookei BUVID3
	urlConfig["CSRF"] = "897dcccd78dfacb98518e8db9328949d"						//b站cookei CSRF
	urlConfig["BILIBILI_UID"] = "1268950779"									//b站cookei BILIBILI_UID


	urlConfig["LOGIN"] = "http://passport.bilibili.com/web/login/v2"
	urlConfig["LOGIN_CAPTCHA"] = "https://passport.bilibili.com/web/captcha/combine?plat=6"
	urlConfig["LOGIN_HASH"] = "http://passport.bilibili.com/login?act=getkey"


	urlConfig["ANIO"] = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?visitor_uid=1268950779&host_uid=1335505500&offset_dynamic_id=0&need_top=1&platform=web" //对方空间 visitor_uid是自己的id,host_uid对方uid
	urlConfig["MY"] = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?visitor_uid=1268950779&host_uid=1268950779&offset_dynamic_id=0&need_top=1&platform=web" //自己空间
	urlConfig["UP_MODIFY"] = "https://api.bilibili.com/x/relation/modify"
	urlConfig["UP_REPOST"] = "https://api.vc.bilibili.com/dynamic_repost/v1/dynamic_repost/repost"
	urlConfig["USER_INFO"] = "http://api.bilibili.com/x/web-interface/nav"
	urlConfig["REPLY"] = "https://api.bilibili.com/x/v2/reply/add"
	urlConfig["ORDINARY"] = "https://api.vc.bilibili.com/dynamic_svr/v2/dynamic_svr/create" //发布动态
	urlConfig["ANIO_DEL"] = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/rm_dynamic" //删除

	urlConfig["VIDEO_REPOST"] = "https://api.vc.bilibili.com/dynamic_repost/v1/dynamic_repost/share"

	urlConfig["LIKE_DYNAMIC"] = "https://api.vc.bilibili.com/dynamic_like/v1/dynamic_like/thumb"
	urlConfig["LIKE_VIDEO"] = "https://api.bilibili.com/x/web-interface/archive/like"


	//urlConfig["HUDONG_HOT"] = "https://api.vc.bilibili.com/topic_svr/v1/topic_svr/topic_new?topic_id=3230836"
	//urlConfig["HUDONG_NEW"] = "https://api.vc.bilibili.com/topic_svr/v1/topic_svr/topic_history?topic_name=%E4%BA%92%E5%8A%A8%E6%8A%BD%E5%A5%96&offset_dynamic_id=" //484787016136660229



	urlConfig["DO_YIFARUHUN"] = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/get_dynamic_detail?dynamic_id=532883838027389857"  //来自转发抽奖娘的动态id
	urlConfig["DO_CHOUJIANG"] = "https://api.bilibili.com/x/activity/lottery/do"  //抽奖链接
	urlConfig["DO_CHOUJIANGNUM"] = "https://api.bilibili.com/x/activity/lottery/mytimes?sid="  //查询抽奖次数
	urlConfig["DO_CHOUJIANGNUMADD"] = "https://api.bilibili.com/x/activity/lottery/addtimes"  //增加抽奖次数



	//拜年活动
	urlConfig["DO_BNJ"] = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/get_dynamic_detail?dynamic_id=489463763067784879"  //485011372337648035最新 483153463975687658一发入魂动态id





	return urlConfig
}
