package webadmins

import (
	"github.com/gin-gonic/gin"
	"my_gin/middleware/session"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/logger"
	"time"
)

type User struct {
}


func NewUser() *User {
	return &User{}
}


func (this *User) Code(ctx *gin.Context){
	modelCap := models.NewModelCaptcha()
	modelCap.Captcha(ctx, 4)
	return
}
func (this *User) Login(ctx *gin.Context){
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

	sessionData, ok := ctx.Get(session.SessionDataName)
	if !ok {
		panic("session middleware")
	}
	sd := sessionData.(models.MemSessionData)
	expire := time.Now().Unix() + global.WEB_ADMINS_LOGIN_EXPIRE
	sd.UpdateExpire(expire)
	sd.Set("expire", expire)
	sd.Save()


	logger.Info("user", "[Login]", "sessionData=", models.RedisMgr.SessionData)
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}

func (this *User) Logout(ctx *gin.Context){
	data := make(map[string]interface{})
	logger.Info("GetAuth", data)
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}

func (this *User) GetSessionInfo(ctx *gin.Context){
	data := make(map[string]interface{})
	sessionId := ctx.PostForm("session")
	sd, ok := models.RedisMgr.SessionData[sessionId]
	if !ok {
		global.JsonRet(ctx, global.ERROR, "数据为空", data)
		return
	}
	data["sessionId"] = sessionId
	data["sessionInfo"] = sd
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}