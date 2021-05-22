package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("can not read config file: ", err.Error())
		return
	}
}
