package routers

import (
	v1 "server-manager/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/users", v1.CreateUser)
	r.DELETE("/users/:userId", v1.DeleteUser)
	r.GET("/users", v1.GetUsers)
	r.GET("/users/:userId", v1.GetUser)

	return r
}
