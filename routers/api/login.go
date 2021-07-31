package api

import (
	"net/http"
	"server-manager/pkg/utils"
	"server-manager/service"

	"github.com/cloudflare/cfssl/log"
	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	var requestData service.UserLoginService
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error("解析失败：", err.Error())
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求失败，请重试！")
		return
	}
	responseData, err := requestData.Login()
	if err != nil {
		log.Error("登录失败：", err.Error())
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求失败，请重试！")
		return
	}
	utils.HandleSuccessResponse(c, http.StatusOK, responseData, "")
}

// 登出
func Logout(c *gin.Context) {

}
