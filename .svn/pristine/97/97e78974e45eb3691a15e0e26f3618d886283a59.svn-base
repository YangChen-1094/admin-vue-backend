package webadmins

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/logger"
)

type deploy struct {
}
func NewDeploy() *deploy{
	return &deploy{}
}

func (this *deploy) Code(ctx *gin.Context){
	modelCap := models.NewModelCaptcha()
	modelCap.Captcha(ctx, 4)
	return
}

func (this *deploy) CodeVerify(ctx *gin.Context){
	data := make(map[string]interface{})
	modelCap := models.NewModelCaptcha()
	code := ctx.PostForm("code")
	check := modelCap.CaptchaVerify(ctx, code)
	if check {
		global.JsonRet(ctx, global.SUCCESS, "", data)
	} else {
		global.JsonRet(ctx, global.INVALID_PARAMS, "", data)
	}
	return
}

func (this *deploy) SessionSet(ctx *gin.Context){
	session := sessions.Default(ctx)
	session.Set("session_01", "123123")
	session.Save()
	data := make(map[string]interface{})
	data["session"] = session
	logger.Info("test", "set session", session)
	global.JsonRet(ctx, global.SUCCESS, "", data)
}
func (this *deploy) SessionGet(ctx *gin.Context){
	session := sessions.Default(ctx)
	value := session.Get("session_01")
	data := make(map[string]interface{})
	data["value"] = value
	data["session"] = session
	logger.Info("test", "get session", session)
	global.JsonRet(ctx, global.SUCCESS, "", data)
}
