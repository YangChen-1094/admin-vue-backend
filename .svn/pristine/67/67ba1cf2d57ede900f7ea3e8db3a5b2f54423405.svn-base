package models

import (
	"time"
)

type ModelAuth struct {
}

func (this *ModelAuth) CheckAuth(username, password string) (int,string) {
	var auth Auth
	data := Auth{Username: username, Password: password}
	db.Select([]string{"id","token"}).Where(data).First(&auth)
	if auth.ID > 0 {
		return auth.ID, auth.Token
	}
	return 0,""
}

//获取这个用户上次的token
func (this *ModelAuth) GetDbToken(id int) string {
	var auth Auth
	db.Select("token").Where("id=?", id).First(&auth)
	if auth.Token != "" {
		return auth.Token
	}
	return ""
}


//更新用户的token值
func (this *ModelAuth) UpdateAuth(id int, token string) bool {
	wh := make(map[string]interface{})
	wh["id"] = id
	db.Model(&Auth{}).Where(wh).Updates(map[string]interface{}{"time" : time.Now().Unix(), "token" : token})
	return true
}