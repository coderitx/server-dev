package utils

import (
	"blog-server/global"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type JWTPayload struct {
	NickName string `json:"nick_name"`
	Role     int    `json:"role"`
	UserID   uint   `json:"user_id"`
}

type CustomClaims struct {
	JWTPayload
	jwt.StandardClaims
}

// GenerateToken 创建token
func GenerateToken(user JWTPayload) (string, error) {
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.GlobalC.JWT.Expires))),
			Issuer:    global.GlobalC.JWT.IsUser,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(global.GlobalC.JWT.Secret))
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GlobalC.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
