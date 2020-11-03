package routers

import (
	"github.com/gin-gonic/gin"
	"my_gin/api/webadmins"
	_ "my_gin/docs" //swagger文档必须要有这个
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
	system := router.Group("/webadmins/system")
	{
		api := webadmins.NewSystem()
		system.POST("/apitest", api.ApiTest)
	}
}

//用户相关
func initUserRouter(router *gin.Engine){
	user := router.Group("/webadmins/user")
	{
		api := webadmins.NewUser()
		user.GET("/vcode", api.Code)
		user.POST("/login", api.Login)
		user.POST("/logout", api.Logout)
	}
}

//综合的接口
func initDeployRouter(router *gin.Engine){
	deploy := router.Group("/webadmins/deploy")
	{
		api := webadmins.NewDeploy()
		deploy.POST("/vcode/verify", api.CodeVerify)
		deploy.GET("/set", api.SessionSet)
		deploy.POST("/get", api.SessionGet)
	}
}

