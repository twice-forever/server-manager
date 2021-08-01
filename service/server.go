package service

import (
	"errors"
	"server-manager/models"
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
		return errors.New("保存数据库失败")
	}

	return nil
}

// 心跳检测服务
func Heartbeat() {

}

// 服务详情
func ServerDetail() {

}
