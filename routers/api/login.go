package api

import (
	"net/http"
	"server-manager/pkg/utils"
	"server-manager/service"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	var requestData service.UserLoginService
	if err := c.ShouldBind(&requestData); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求失败，请重试！")
		return
	}
	responseData := requestData.Login(c)
	utils.HandleSuccessResponse(c, http.StatusOK, responseData, "")
}

// 登出
func Logout(c *gin.Context) {

}
