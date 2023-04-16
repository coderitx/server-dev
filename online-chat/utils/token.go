package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const JwtKey = "online-chat"

type UserClaim struct {
	Id       uint
	Identity string
	Name     string
	Phone    string
	Email    string
	jwt.StandardClaims
}

// GenerateToken
// 根据用户信息生成token
func GenerateToken(id uint, identity, name, phone, email string, expiredAt int) (string, error) {
	uc := UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expiredAt)).Unix(),
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	token, err := claims.SignedString([]byte(JwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func AnalyzeToken(token string) (*UserClaim, error) {
	uc := new(UserClaim)
	var err error
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		if claims != nil {
			// 过期时间戳
			expireTime := time.Now().Unix() - uc.ExpiresAt
			if int(expireTime) < 300 {
				return uc, errors.New("402")
			}
		}
		return nil, err
	}

	if !claims.Valid {
		ve, _ := err.(*jwt.ValidationError)
		if ve.Errors == jwt.ValidationErrorExpired {
			return uc, errors.New("token expired")
		}
		return uc, errors.New("token is invalid")
	}
	return uc, err
}
