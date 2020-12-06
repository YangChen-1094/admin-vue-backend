package jwt

import (
	"github.com/gin-gonic/gin"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/logger"
	"my_gin/pkg/util"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlPath := ctx.Request.URL.Path
		if urlPath == "/webadmins/user/logout" {
			ctx.SetCookie("token", "", -1, "/", "localhost", false, false)
			global.JsonRet(ctx, 200, "logout success", nil)
			ctx.Abort()
			return
		}

		var aNotCheckUrl = [...]string{
			"/v1/auth/getToken",
			"/webadmins/user/vcode",
			"/webadmins/user/captchaId",
			"/webadmins/user/codeImg",
			"/webadmins/user/login",
		}

		if util.InArray(urlPath, aNotCheckUrl) {//指定url 不需要通过 token验证
			ctx.Next()
			return
		}

		var code int
		code = global.SUCCESS
		token := ctx.GetHeader("token")
		token = strings.Trim(token, "\b")
		if token == "" {
			code = global.LOGIN_ERROR
		} else {
			tokenParam, err := util.ParseToken(token)
			if tokenParam == nil || err != nil { //验证token失败
				code = global.ERROR_AUTH_CHECK_TOKEN_FAIL
				global.JsonRet(ctx,  global.LOGIN_ERROR, "token已失效！", nil)
				ctx.Abort()
				return
			} else if time.Now().Unix() > tokenParam.ExpiresAt { //token 过期
				code = global.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
			logger.Info("user", "[Login]", ", tokenParam=",tokenParam)
			ModelAuth := models.NewModelAuth()
			dbToken := ModelAuth.GetDbToken(tokenParam.Id)//检查这个用户是否存在
			if dbToken != token {
				global.JsonRet(ctx,  global.LOGIN_ERROR, "token已失效！", nil)
				ctx.Abort()
				return
			}
		}

		if code != global.SUCCESS { //token 验证不通过
			global.JsonRet(ctx, code, "", nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
