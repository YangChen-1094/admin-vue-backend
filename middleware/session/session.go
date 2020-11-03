package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"my_gin/models"
)

const sessionMaxAge = 3600
const sessionSecret = "mhjy"
const SessionIdName = "session_id"		//存放sessionId的字段名
const SessionDataName = "session_data"	//存放session数据的字段名

type Code struct {
	Code string
}

func LoginSessionMiddleware() gin.HandlerFunc {
	models.RedisMgr = models.NewRedisManager()
	models.RedisMgr.Init("127.0.0.1:6379")
	return func(ctx *gin.Context) {
		sessionId, err := ctx.Cookie(SessionIdName)
		var sessionData models.MemSessionData
		if err != nil {//没传session
			sessionData = models.RedisMgr.CreateSessionData()
			sessionId = sessionData.GetId()
		}else{//传了sessionId
			sessionData, err = models.RedisMgr.GetSessionData(sessionId)
			if err != nil {//内存或者redis 没有对应session
				sessionData = models.RedisMgr.CreateSessionData()
			}
			sessionId = sessionData.GetId()
		}
		ctx.Set(SessionDataName, sessionData) //存放这个session数据， 为后面使用
		ctx.SetCookie(SessionIdName, sessionId, 1200, "/", "localhost", false, true)
		ctx.Next()
	}
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
