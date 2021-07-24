package setting

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	HTTPPort int

	JwtSecret string
)

func Setup() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("can not read config file: ", err.Error())
		return
	}

	loadApp()
}

func loadApp() {
	JwtSecret = viper.GetString("app.jwt-secret")
}
