package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/util"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = global.SUCCESS
		auth_id := com.StrTo(ctx.PostForm("auth_id")).MustInt()
		token := ctx.PostForm("token")
		token = strings.Trim(token, "\b")
		ctx.Set("test", "111")
		if token == "" || auth_id <= 0 {
			code = global.INVALID_PARAMS
		} else {
			tokenParam, err := util.ParseToken(token)
			if err != nil { //验证token失败
				code = global.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > tokenParam.ExpiresAt { //token 过期
				code = global.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
			ModelAuth := models.NewModelAuth()
			dbToken := ModelAuth.GetDbToken(auth_id)//检查这个用户是否存在
			if dbToken != token {
				global.JsonRet(ctx,  global.ERROR_AUTH_CHECK_TOKEN_FAIL, "token已失效！", nil)
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
