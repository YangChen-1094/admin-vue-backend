package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const sessionMaxAge = 3600
const sessionSecret = "mhjy"

type Code struct {
	Code string
}
// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := getSessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func getSessionConfig() sessions.Store {
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}

func Cookie(name string) gin.HandlerFunc {
	store := getCookieConfig()
	return sessions.Sessions(name, store)
}

func getCookieConfig() sessions.Store {
	var store sessions.Store
	store = cookie.NewStore([]byte("cookie"))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}
