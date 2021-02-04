package config

func GetUrlConfig() map[string]interface{} {
	// 初始化数据库配置map
	urlConfig := make(map[string]interface{})

	urlConfig["HOST"] = "https://www.laosiji.com/"

	urlConfig["CAR_BRAND"] = urlConfig["HOST"].(string)+"auto/"

	urlConfig["CAR_MODEL"] = urlConfig["HOST"].(string)+"api/sellingcarslist"
	urlConfig["CAR_SERIES_CONFIG"] = urlConfig["HOST"].(string)+"api/car/carsconfig"



	return urlConfig
}
