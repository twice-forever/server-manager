package models

import (
	"time"

	"gorm.io/gorm"
)

// 用户结构
type User struct {
	ID        int
	Username  string
	Password  string `json:"-"`
	PWDSalt   string `json:"-"`
	RealName  string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}

// 获取用户
func (u *User) GetUser() error {
	if err := DB.Where(&User{Username: u.Username}).First(u).Error; err != nil {
		return err
	}
	return nil
}

// 增加用户
func (u *User) AddUser() error {
	if err := DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// 删除用户
func (u *User) DeleteUser() error {
	if err := DB.Delete(u).Error; err != nil {
		return err
	}
	return nil
}

// 更新用户
func (u *User) UpdateUser() error {
	if err := DB.Model(u).Select("username", "real_name", "avatar_url").Updates(u).Error; err != nil {
		return err
	}
	return nil
}

// 更新密码
func (u *User) UpdateUserPassword() error {
	if err := DB.Model(u).Select("password").Updates(u).Error; err != nil {
		return err
	}
	return nil
}
