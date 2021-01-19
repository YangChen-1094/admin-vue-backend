package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "my_gin/docs" //swagger文档必须要有这个
	"my_gin/middleware/cors"
	"my_gin/middleware/session"
	"my_gin/pkg/setting"
)

func InitRouter() *gin.Engine{
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Cors())//跨域
	router.Use(session.Session("mhjy"))
	//router.Use(session.LoginSessionMiddleware())//登录 不使用session
	gin.SetMode(setting.DeployConfig.Server.RunMode)
	router.GET("/Swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//router.Use(jwt.JWT())//验证token中间件

	//业务api路由
	initApiRouters(router)

	//GM后台路由
	initWebAdminsRouters(router)

	return router
}
