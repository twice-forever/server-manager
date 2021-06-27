package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	PWDSalt   string         `json:"pwdSalt"`
	RealName  string         `json:"realName"`
	AvatarURL string         `json:"avatarUrl"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type ShowUser struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	RealName  string    `json:"realName"`
	AvatarURL string    `json:"avatarUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ChangePassword struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
