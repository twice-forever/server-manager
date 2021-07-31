package utils

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"

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

// 密码生成
func GetRandomString(len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result []byte = make([]byte, len)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		result[i] = str[r.Intn(64)]
	}
	return string(result)
}
