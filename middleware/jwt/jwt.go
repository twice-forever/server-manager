package jwt

import (
	"net/http"
	"server-manager/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// 登录信息检查中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "登录信息为空")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求头有误")
			c.Abort()
			return
		}

		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无效的token")
			c.Abort()
			return
		}

		c.Set("userId", mc.Id)
		c.Next()
	}
}
