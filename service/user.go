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
	Username string `binding:"required"`
	Password string `binding:"required,min=4,max=64"`
	PWDSalt  string
}

// Login 用户登录函数
func (service *UserLoginService) Login() (map[string]interface{}, error) {
	response := map[string]interface{}{}
	if err := CheckUser(service); err != nil {
		return response, err
	}

	token, err := utils.GenerateToken(service.ID)
	if err != nil {
		return response, err
	}

	response["token"] = token
	return response, nil
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
	if hashPassword == checkUser.Password {
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

// 创建admin用户
func CreateAdmin() {
	admin := models.User{
		Username: "admin",
		RealName: "admin",
	}

	if err := admin.GetUser(); err == nil {
		return
	}

	password := utils.GetRandomString(16)
	salt := utils.GenerateSalt()
	hashPassword := utils.GeneratHashPWD(password, salt)

	admin.PWDSalt = salt
	admin.Password = hashPassword

	if err := admin.AddUser(); err != nil {
		log.Error("增加admin用户失败：", err.Error())
		return
	}

	log.Info("username: admin")
	log.Info("password: ", password)
}
