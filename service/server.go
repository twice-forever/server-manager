package service

import (
	"errors"
	"server-manager/models"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
	Desc string `json:"desc"`
}

func (s Server) RegisterServer() error {
	server := &models.Server{
		ID:   s.ID,
		Name: s.Name,
		IP:   s.IP,
		Port: s.Port,
		Desc: s.Desc,
	}

	// 检查重名
	if err := server.GetByName(); err == nil {
		return errors.New("存在重名服务")
	}

	// 检查重复ip+port
	if err := server.GetByIP(); err == nil {
		return errors.New("存在重复ip+port")
	}

	// 写入数据库
	if err := server.Create(); err != nil {
		return err
	}

	return nil
}

func (s Server) Delete() error {
	server := &models.Server{
		ID:   s.ID,
		Name: s.Name,
		IP:   s.IP,
		Port: s.Port,
		Desc: s.Desc,
	}

	if err := server.Delete(); err != nil {
		return err
	}
	return nil
}

func GetServers(c *gin.Context) ([]models.Server, error) {
	servers := make([]models.Server, 0, 10)
	tempDB := models.DB.Scopes(Paginate(c)).Model(models.User{})
	if err := tempDB.Find(&servers).Error; err != nil {
		return servers, err
	}
	if err := tempDB.Find(&servers).Error; err != nil {
		return servers, err
	}
	return servers, nil
}

// 心跳检测服务
func Heartbeat() {

}

// 服务详情
func ServerDetail() {

}
