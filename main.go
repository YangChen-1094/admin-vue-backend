package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"log"
	"my_gin/models"
	modelGrpc "my_gin/models/grpc"
	"my_gin/pkg/setting"
	"my_gin/routers"
	"net/http"
)

func init() {//初始化
	setting.Setup()
	models.Setup()
}

func main() {
	router := routers.InitRouter()
	router.Use(static.Serve("/", static.LocalFile("./dist/", false)))
	//router.Static("./static", "/static")
	//router.GET("/", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "index.html", nil)
	//	return
	//})
	//endless.DefaultReadTimeOut = setting.ReadTimeOut
	//endless.DefaultWriteTimeOut = setting.WriteTimeOut
	//endless.DefaultMaxHeaderBytes = 1 << 20 //请求头的最大字节数 2^20 即：1024 * 1024b => 1M
	//ipPort := fmt.Sprintf(":%d", setting.HttpPort)

	//server := endless.NewServer(ipPort, router)//windows系统没有 syscall.SIGUSR1 和 syscall.SIGUSR2，所以会报此错
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())//启动时输出pid
	//}
	modelGrpc.Register()
	Mysql := setting.Mysql.GetMysqlClient("wbAdminDatas", "")
	data := models.ItemType{}
	errs := Mysql.Where("id=?", 12).Find(&data).Error
	fmt.Println("Mysql wbAdminDatas:", Mysql)
	fmt.Println("Mysql data:", data)
	fmt.Println("errs:", errs)

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", setting.DeployConfig.Server.HttpIp,setting.DeployConfig.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.DeployConfig.Server.ReadTimeOut,
		WriteTimeout:   setting.DeployConfig.Server.WriteTimeOut,
		MaxHeaderBytes: 1 << 20, //请求头的最大字节数 2^20 即：1024 * 1024b => 1M
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
