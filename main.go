package main

import (
	"server-manager/models"
	"server-manager/pkg/setting"
	"server-manager/routers"
)

func init() {
	// 读取配置文件
	setting.Setup()

	// 连接数据库
	models.Setup()
}

func main() {
	// 初始化路由并运行
	r := routers.InitRouter()
	r.Run()
}
