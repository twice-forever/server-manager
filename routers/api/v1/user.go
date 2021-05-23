package v1

import (
	"net/http"
	"server-manager/models"
	"server-manager/pkg/utils"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 创建用户
func CreateUser(c *gin.Context) {
	// 解析请求数据
	var requestData models.User
	if err := c.ShouldBind(&requestData); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求失败，请重试！")
		return
	}

	// 表单验证
	valid := validation.Validation{}
	valid.Required(requestData.Username, "username").Message("username")
	valid.Required(requestData.Password, "password").Message("password")

	// 增加数据到数据库
	requestData.PWDSalt = utils.GenerateSalt()
	requestData.Password = utils.GeneratHashPWD(requestData.Password, requestData.PWDSalt)
	result := utils.ConnectDB.Create(&requestData)
	if result.Error != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "增加用户失败，请重试！")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, gin.H{"id": requestData.ID}, "")
}
