package respdata


type BilibiliCode struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Ttl int `json:"ttl"`
}

type BilibiliChouJiangNumData struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Ttl int `json:"ttl"`
	Data struct{
		Times int `json:"times"`
	}`json:"data"`
}

type BilibiliCaptcha struct {
	Data struct{
		Result struct{
			Gt string `json:"gt"`
			Challenger string `json:"challenger"`
			Key string `json:"key"`
		} `json:"result"`
	} `json:"data"`
}
type BilibiliHash struct {
	Hash string `json:"hash"`
	Key string `json:"key"`
}


type BilibiliSpaceHistory struct {
	Data struct{
		Cards []struct{
			Card string `json:"card"`
			Desc struct{
				OrigDyIdStr string `json:"orig_dy_id_str"`  //父内容
				Origin struct{
					DynamicIdStr string `json:"dynamic_id_str"`
					RidStr string `json:"rid_str"`   //目标评论区id
					Type int `json:"type"`
					Uid int `json:"uid"`		//对方uid
				} `json:"origin"`
				PreDyIdStr string `json:"pre_dy_id_str"`  //本体转发文字动态id
				Previous struct{
					DynamicIdStr string `json:"dynamic_id_str"` //子动态id
					RidStr string `json:"rid_str"`   //子评论区id
					Type int `json:"type"`     //子动态类型
					Uid int `json:"uid"`    //对方uid
				}`json:"previous"`
			} `json:"desc"`
		} `json:"cards"`

	} `json:"data"`
	JsonData string
}

type BilibiliDynamicDetail struct {
	Code int `json:"code"`
	Data struct{
		Card struct{
			Card string `json:"card"`
			CardJson struct{
				Item struct{
					Description string `json:"description"`
				}`json:"item"`
			}
		} `json:"card"`

	} `json:"data"`
}

type BilibiliActivity struct {
	BaseInfo struct{
		Title string `json:"title"`
		Description string `json:"description"`
		Keywords string `json:"keywords"`
		SharePicture string `json:"sharePicture"`
	}`json:"BaseInfo"`
	LotteryNew []struct{
		LotteryId string `json:"lotteryId"`
	}`json:"h5-lottery-new"`
	FollowNew []struct{
		Uid string `json:"uid"`
	}`json:"h5-follow-new"`
	PcLotteryNew []struct{
		LotteryId string `json:"lotteryId"`
	}`json:"pc-lottery-new"`
}




type BilibiliSpaceHistoryCardJson struct {
	Item struct{
		Content string `json:"content"`
	}
	OriginUser struct{
		Info struct{
			Uid int64 `json:"uid"`
			Uname string `json:"uname"`
		} `json:"info"`
	} `json:"origin_user"`
}

type BilibiliUserInfo struct {
	Mid string `json:"mid"`
	Uname string `json:"uname"`
}


type BilibiliDo struct {
	Code int `json:"code"`
	Data []struct{
		GiftName string `json:"gift_name"`
	} `json:"data"`
	Message string `json:"message"`
}