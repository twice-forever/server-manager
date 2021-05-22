package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ConnectDB *gorm.DB

func ConnectDatabase() {
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	ip := viper.GetString("database.ip")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, dbname)
	ConnectDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("can not connect database: ", err.Error())
		return
	}
	log.Info(ConnectDB)
}
