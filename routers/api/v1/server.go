package v1

import (
	"net/http"
	"server-manager/pkg/utils"
	"server-manager/service"
	"strconv"

	"github.com/cloudflare/cfssl/log"
	"github.com/gin-gonic/gin"
)

// 服务注册
func RegisterServer(c *gin.Context) {
	// 解析请求数据
	var requestData service.Server
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error("数据解析失败：", err.Error())
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "数据解析失败，请重试！")
		return
	}

	// 注册服务
	if err := requestData.RegisterServer(); err != nil {
		log.Error("服务注册失败：", err.Error())
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "服务注册失败，请重试！")
		return
	}

	utils.HandleSuccessResponse(c, http.StatusOK, &gin.H{"ID": requestData.ID}, "")
}

// 删除服务
func DeleteServer(c *gin.Context) {
	// 解析请求数据
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无服务id")
		return
	}
	requestData := service.Server{
		ID: id,
	}

	// 删除数据
	if err := requestData.Delete(); err != nil {
		log.Error("服务删除失败：", err.Error())
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "服务删除失败，请重试！")
		return
	}

	utils.HandleSuccessResponse(c, http.StatusOK, &gin.H{"ID": requestData.ID}, "")
}

// 获取服务列表
func GetServerList(c *gin.Context) {

}

// 获取服务详情
func GetServer(c *gin.Context) {

}

// 更新服务
func UpdateServer(c *gin.Context) {

}

// 启动服务
func StartServer(c *gin.Context) {

}

// 停止服务
func StopServer(c *gin.Context) {

}
