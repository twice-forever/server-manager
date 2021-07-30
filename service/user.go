package service

import (
	"errors"
	"net/http"
	"server-manager/models"
	"server-manager/pkg/utils"

	"github.com/cloudflare/cfssl/log"
	"github.com/gin-gonic/gin"
)

// 修改密码
type ChangePassword struct {
	Password    string
	NewPassword string
}

// 管理用户登录的服务
type UserLoginService struct {
	ID       int
	Username string `binding:"required,email"`
	Password string `binding:"required,min=4,max=64"`
	PWDSalt  string
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) map[string]interface{} {
	response := map[string]interface{}{}
	if err := CheckUser(service); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "登录失败，请重试！")
		return response
	}

	token, err := utils.GenerateToken(service.ID)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "登录失败，请重试！")
		return response
	}

	response["token"] = token
	return response
}

// 查询用户名称和确定密码正确
func CheckUser(user *UserLoginService) error {
	password := user.Password
	checkUser := &models.User{
		Username: user.Username,
	}
	if err := checkUser.GetUser(); err != nil {
		log.Error("can not connect database: ", err.Error())
		return err
	}

	hashPassword := utils.GeneratHashPWD(password, checkUser.PWDSalt)
	if hashPassword == user.Password {
		return nil
	}

	return errors.New("密码不正确")
}

// 获取用户
func GetUsers(c *gin.Context) ([]models.User, error) {
	users := make([]models.User, 0, 10)
	tempDB := models.DB.Scopes(Paginate(c)).Model(models.User{})
	if err := tempDB.Find(&users).Error; err != nil {
		return users, err
	}
	if err := tempDB.Find(&users).Error; err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "查询用户错误")
		return users, err
	}
	return users, nil
}
