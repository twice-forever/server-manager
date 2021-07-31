package main

import (
	"server-manager/models"
	"server-manager/pkg/setting"
	"server-manager/routers"
	"server-manager/service"
)

func init() {
	// 读取配置文件
	setting.Setup()

	// 连接数据库
	models.Setup()

	// 初始化用户
	service.CreateAdmin()
}

func main() {
	// 初始化路由并运行
	r := routers.InitRouter()
	r.Run()
}
