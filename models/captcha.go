package models

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"my_gin/pkg/logger"
	"net/http"
	"time"
)

//类
type ModelCaptcha struct {
}

func (this *ModelCaptcha)Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 120, 50
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	logger.Info("GetAuth", "set captchaId=", captchaId, ",session=",session)
	_ = session.Save()
	_ = this.Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
func (this *ModelCaptcha)CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	captchaId := session.Get("captcha")
	logger.Info("GetAuth", "code=", code, ",get captchaId=", captchaId, ",session=", session)
	if captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func (this *ModelCaptcha)Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	cookie := &http.Cookie{
		Name: "captcha",
		Value: id,
		Path: "/",
		MaxAge: int(time.Hour * 24 / time.Second) ,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: false,
		Expires: time.Date(2020, 11, 29, 15, 30, 10, 0, time.Local),
	}
	var content bytes.Buffer
	//w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")//
	//w.Header().Set("Pragma", "no-cache")
	//w.Header().Set("Expires", "0")
	w.Header().Add("Set-Cookie", cookie.String())
	//http.SetCookie(w, cookie)
	//session.Cookie("captcha")
	switch ext {
		case ".png":
			w.Header().Set("Content-Type", "image/png")
			_ = captcha.WriteImage(&content, id, width, height)
		case ".wav":
			w.Header().Set("Content-Type", "audio/x-wav")
			_ = captcha.WriteAudio(&content, id, lang)
		default:
			return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
