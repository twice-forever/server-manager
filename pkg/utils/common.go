package utils

import (
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
)

// 生成盐
func GenerateSalt() string {
	uuid := uuid.NewString()
	hash := sha256.Sum256([]byte(uuid))
	return fmt.Sprintf("%x", hash)
}

// 生成hash密码
func GeneratHashPWD(password, salt string) string {
	fullString := fmt.Sprintf("%s%s", password, salt)
	hash := sha256.Sum256([]byte(fullString))
	return fmt.Sprintf("%x", hash)
}
