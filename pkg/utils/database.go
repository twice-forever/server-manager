package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
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

	var err error
	ConnectDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("can not connect database: ", err.Error())
		return
	}
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
