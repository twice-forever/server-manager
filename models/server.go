package models

import "gorm.io/gorm"

type Server struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
	Desc string `json:"desc"`

	gorm.Model
}

func (s *Server) Get() error {
	if err := DB.First(s).Error; err != nil {
		return err
	}
	return nil
}

func (s Server) Delete() error {
	if err := DB.Delete(&s).Error; err != nil {
		return err
	}
	return nil
}

func (s *Server) Create() error {
	if err := DB.Create(s).Error; err != nil {
		return err
	}
	return nil
}

func (s *Server) Update() error {
	if err := DB.Model(s).Updates(*s).Error; err != nil {
		return err
	}
	return nil
}
