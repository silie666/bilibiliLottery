package config

import (
	"github.com/spf13/viper"
	"log"
)

var Env *viper.Viper

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath("/files/go/bilibiliLottery/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	Env = viper.GetViper()
}
