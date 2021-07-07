package api

import (
	"net/http"
	"server-manager/models"
	"server-manager/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	var requestData models.User
	if err := c.ShouldBind(&requestData); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求失败，请重试！")
		return
	}

	if err := utils.CheckUser(&requestData); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "登录失败，请重试！")
		return
	}

	token, err := utils.GenerateToken(requestData.ID)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "登录失败，请重试！")
		return
	}

	utils.HandleSuccessResponse(c, http.StatusOK, gin.H{
		"token": token,
	}, "")
}

// 登出
func Logout(c *gin.Context) {

}
