package routers

import (
	v1 "server-manager/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.POST("/users", v1.CreateUser)
	r.DELETE("/users/:userId", v1.DeleteUser)

	return r
}
