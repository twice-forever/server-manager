package main

import (
	"server-manager/pkg/utils"
	"server-manager/routers"
)

func main() {
	// 读取配置文件
	utils.ReadConfig()

	// 连接数据库
	utils.ConnectDatabase()

	// 初始化路由并运行
	r := routers.InitRouter()
	r.Run()
}
