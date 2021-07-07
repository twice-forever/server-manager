package utils

import (
	"errors"
	"server-manager/models"

	log "github.com/sirupsen/logrus"
)

func CheckUser(user *models.User) error {
	password := user.Password
	if err := ConnectDB.Where(&models.User{Username: user.Username}).First(&user).Error; err != nil {
		log.Error("can not connect database: ", err.Error())
		return err
	}

	hashPassword := GeneratHashPWD(password, user.PWDSalt)
	if hashPassword == user.Password {
		return nil
	}

	return errors.New("密码不正确")
}
