package routers

import (
	"github.com/gin-gonic/gin"
	"my_gin/api/webadmins"
	_ "my_gin/docs" //swagger文档必须要有这个
	"my_gin/middleware/jwt"
)

/**
 * 此处都是GM后台的接口路由
 */

func initWebAdminsRouters(router *gin.Engine){
	initSystemRouter(router)
	initUserRouter(router)
	initDeployRouter(router)
}

//系统接口
func initSystemRouter(router *gin.Engine){
	system := router.Group("/webadmins/system").Use(jwt.JWT())//验证token中间件
	{
		api := webadmins.NewSystem()
		system.POST("/apitest", api.ApiTest)
	}
}

//用户相关
func initUserRouter(router *gin.Engine){
	user := router.Group("/webadmins/user").Use(jwt.JWT())//验证token中间件
	{
		api := webadmins.NewUser()
		user.GET("/vcode", api.Code)
		user.POST("/login", api.Login)
		user.POST("/logout", api.Logout)
		user.POST("/getSessionInfo", api.GetSessionInfo)
	}
}

//综合的接口
func initDeployRouter(router *gin.Engine){
	deploy := router.Group("/webadmins/deploy").Use(jwt.JWT())//验证token中间件
	{
		api := webadmins.NewDeploy()
		deploy.POST("/vcode/verify", api.CodeVerify)
		deploy.GET("/set", api.SessionSet)
		deploy.POST("/get", api.SessionGet)
	}
}

