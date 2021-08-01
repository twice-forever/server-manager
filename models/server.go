package models

import (
	"time"

	"gorm.io/gorm"
)

type Server struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
	Desc string `json:"desc"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// 根据id查找
func (s *Server) Get() error {
	if err := DB.First(s).Error; err != nil {
		return err
	}
	return nil
}

// 根据name查找
func (s *Server) GetByName() error {
	if err := DB.Where("name = ?", s.Name).First(s).Error; err != nil {
		return err
	}
	return nil
}

// 根据ip+port查找
func (s *Server) GetByIP() error {
	if err := DB.Where(&Server{IP: s.IP, Port: s.Port}).First(s).Error; err != nil {
		return err
	}
	return nil
}

// 删除
func (s Server) Delete() error {
	if err := DB.Delete(&s).Error; err != nil {
		return err
	}
	return nil
}

// 创建
func (s *Server) Create() error {
	if err := DB.Create(s).Error; err != nil {
		return err
	}
	return nil
}

// 更新
func (s *Server) Update() error {
	if err := DB.Model(s).Updates(*s).Error; err != nil {
		return err
	}
	return nil
}
