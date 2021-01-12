package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Params struct {//定义的jwt里面存啥
	Id       int    `json:"id"`
	Username string `json:"username"`
	//Password string `json:"password"`//不能有私密信息
	jwt.StandardClaims
}

func GenerateToken(id int, username string, jwtSecret []byte) (string, error) {
	expire := time.Now().Add(24 * time.Hour)
	param := Params{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenParam := jwt.NewWithClaims(jwt.SigningMethodHS256, param)
	token, err := tokenParam.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string, jwtSecret []byte) (*Params, error) {
	tokenParam, err := jwt.ParseWithClaims(token, &Params{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenParam != nil {
		if param, ok := tokenParam.Claims.(*Params); ok && tokenParam.Valid {
			return param, nil
		}
	}
	return nil, err
}
