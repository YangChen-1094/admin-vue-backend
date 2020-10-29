package webadmins

import (
	"github.com/gin-gonic/gin"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/logger"
)

type user struct {
}


func NewUser() *user{
	return &user{}
}


func (this *user) Code(ctx *gin.Context){
	modelCap := models.NewModelCaptcha()
	modelCap.Captcha(ctx, 4)
	return
}
func (this *user) Login(ctx *gin.Context){
	data := make(map[string]interface{})
	logger.Info("GetAuth", data)
	modelCap := models.NewModelCaptcha()
	code := ctx.PostForm("vcode")
	check := modelCap.CaptchaVerify(ctx, code)
	logger.Info("GetAuth", code, check)
	if !check {
		global.JsonRet(ctx, 401, "验证码错误", data)
		return
	}

	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}

func (this *user) Logout(ctx *gin.Context){
	data := make(map[string]interface{})
	logger.Info("GetAuth", data)
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}