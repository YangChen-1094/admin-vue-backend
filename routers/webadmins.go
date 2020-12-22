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
	//GM后台路由
	webGroup := router.Group("/webadmins")//webadmins路由组
	{
		webGroup.Use(jwt.JWT())
		initUserRouter(webGroup)
		initChannelRouter(webGroup)
		initDeployRouter(webGroup)
		initSystemRouter(webGroup)
	}
}



//用户相关
func initUserRouter(webGroup *gin.RouterGroup){
	user := webGroup.Group("/user")//webadmins 下的user路由组
	{
		api := webadmins.NewUser()
		user.POST("/captchaId", api.CaptchaId)
		user.GET("/codeImg", api.CodeImg)
		user.GET("/vcode", api.Code)
		user.POST("/login", api.Login)
		user.POST("/logout", api.Logout)
		user.POST("/getSessionInfo", api.GetSessionInfo)
	}
}

//综合的接口
func initChannelRouter(webGroup *gin.RouterGroup){
	channel := webGroup.Group("/channel")//webadmins 下的channel路由组
	{
		api := webadmins.NewChannel()
		channel.POST("/list", api.List)
		channel.POST("/modify", api.Modify)
		channel.POST("/add", api.Add)
		channel.POST("/del", api.Delete)
		channel.POST("/import", api.Import)
		channel.POST("/export", api.Export)
	}
}

//综合的接口
func initDeployRouter(webGroup *gin.RouterGroup){
	deploy := webGroup.Group("/deploy")//webadmins 下的deploy路由组
	{
		api := webadmins.NewDeploy()
		deploy.POST("/vcode/verify", api.CodeVerify)
		deploy.GET("/set", api.SessionSet)
		deploy.POST("/get", api.SessionGet)
	}
}

//系统接口
func initSystemRouter(webGroup *gin.RouterGroup){
	system := webGroup.Group("/system")//webadmins 下的system路由组
	{
		api := webadmins.NewSystem()
		system.POST("/apitest", api.ApiTest)
	}
}