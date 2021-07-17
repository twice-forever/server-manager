package utils

import (
	"server-manager/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}

func GenerateToken(userId int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "server-manager",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil && err == nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
