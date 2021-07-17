package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理正确请求
func HandleSuccessResponse(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"data":    data,
		"message": message,
	})
}

// 处理错误请求
func HandleErrorResponse(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    statusCode,
		"data":    data,
		"message": message,
	})
}
