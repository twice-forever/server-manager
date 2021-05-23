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
	AvatarURL string         `json:"avatar_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
