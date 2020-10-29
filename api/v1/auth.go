package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/logger"
	"my_gin/pkg/util"
)

type ApiAuth struct {
}

//参数验证格式
type authValid struct {
	Username string	`valid:"Required; MaxSize(50)"`
	Password string	`valid:"Required; MaxSize(50)"`
}

func NewApiAuth() *ApiAuth{
	return &ApiAuth{}
}

func (this *ApiAuth) GetAuth(ctx *gin.Context){
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	v := validation.Validation{}
	validParam := authValid{Username: username,Password: password}
	ok, _ := v.Valid(validParam)
	if !ok {
		var msg string
		for _, err := range v.Errors {
			msg += err.Key + ":" + err.Message +";"
		}
		global.JsonRet(ctx, global.INVALID_PARAMS, msg, nil)
		return
	}
	ModelAuth := models.NewModelAuth()
	id, _ := ModelAuth.CheckAuth(username, password)//检查这个用户是否存在
	if id <= 0 {
		global.JsonRet(ctx, global.ERROR, "用户不存在", nil)
		return
	}

	token, err := util.GenerateToken(username, password)//生成token
	if err != nil {
		global.JsonRet(ctx, global.ERROR_AUTH_TOKEN, "", nil)
		return
	}
	update := ModelAuth.UpdateAuth(id, token)//更新token到db, 所以一个token是单次有效且有过期时间， 通过v1/getToken接口获取并刷新token，此前的token就不能用了
	if update == false {
		global.JsonRet(ctx, global.ERROR_AUTH_CHECK_TOKEN_FAIL, "更新token失败", nil)
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	logger.Info("GetAuth", data)
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}