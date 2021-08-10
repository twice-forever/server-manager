package routers

import (
	"server-manager/middleware/jwt"
	"server-manager/routers/api"
	v1 "server-manager/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login", api.Login)

	v1Group := r.Group("/v1")
	v1Group.Use(jwt.JWTAuthMiddleware())
	{
		v1Group.POST("/users", v1.CreateUser)
		v1Group.DELETE("/users/:userId", v1.DeleteUser)
		v1Group.GET("/users", v1.GetUsers)
		v1Group.GET("/users/:userId", v1.GetUser)
		v1Group.PUT("/users/:userId", v1.UpdateUsers)
		v1Group.PATCH("/users/:userId/password", v1.ChangePassword)
	}

	serverGroup := r.Group("/servers")
	{
		serverGroup.POST("/register", v1.RegisterServer)
		serverGroup.DELETE("/:id", v1.DeleteServer)
		serverGroup.GET("/", v1.GetServerList)
		serverGroup.GET("/:id", v1.GetServer)
		serverGroup.POST("/", v1.UpdateServer)
		serverGroup.POST("/start/:id", v1.StartServer)
		serverGroup.POST("/stop/:id", v1.StopServer)
	}

	return r
}
