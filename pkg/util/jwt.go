package util

import (
	"github.com/dgrijalva/jwt-go"
	"my_gin/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string)(string, error){
	now := time.Now()
	expire := now.Add(3 * time.Hour)

	param := Params{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer: "gin-blog",
		},
	}

	tokenParam := jwt.NewWithClaims(jwt.SigningMethodHS256, param)
	token, err := tokenParam.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string)(* Params, error){
	tokenParam, err := jwt.ParseWithClaims(token, &Params{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret,nil
	})
	if tokenParam != nil {
		if param, ok := tokenParam.Claims.(*Params); ok && tokenParam.Valid {
			return param,nil
		}
	}
	return nil, err
}