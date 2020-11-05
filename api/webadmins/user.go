package webadmins

import (
	"github.com/gin-gonic/gin"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/logger"
	"my_gin/pkg/util"
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
	if !check {
		global.JsonRet(ctx, 401, "验证码错误", data)
		return
	}

	//不使用session
	//sessionData, ok := ctx.Get(session.SessionDataName)
	//if !ok {
	//	panic("session middleware")
	//}
	//sd := sessionData.(models.MemSessionData)
	//expire := time.Now().Unix() + global.WEB_ADMINS_LOGIN_EXPIRE
	//sd.UpdateExpire(expire)
	//sd.Set("expire", expire)
	//sd.Save()

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")//密码加密
	dePass := util.AesDecrypt(password)

	ModelAuth := models.NewModelAuth()
	id, _ := ModelAuth.CheckAuth(username, dePass)//检查这个用户是否存在
	if id <= 0 {
		global.JsonRet(ctx, global.ERROR, "用户不存在", nil)
		return
	}

	//使用jwt 作为区分每个用户的token
	token, err := util.GenerateToken(id, username)//生成token
	if err != nil {
		global.JsonRet(ctx, global.ERROR, "jwt token生成失败", nil)
		return
	}

	update := ModelAuth.UpdateAuth(id, token)//更新token到db, 所以一个token是单次有效且有过期时间， 通过v1/getToken接口获取并刷新token，此前的token就不能用了
	if update == false {
		global.JsonRet(ctx, global.ERROR_AUTH_CHECK_TOKEN_FAIL, "更新token失败", nil)
		return
	}
	ctx.SetCookie("token", token, 1200, "/", "localhost", false, false)
	logger.Info("user", "[Login]", ", token=",token)
	//并且把这个token保存到cookie，客户端每次带过来

	data["id"] = id
	data["username"] = username
	data["token"] = token
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}

func (this *User) Logout(ctx *gin.Context){
	data := make(map[string]interface{})
	param := ctx.Params
	logger.Info("user", "[Logout]",param)
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