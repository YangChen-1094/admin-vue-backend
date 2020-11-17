package main

import (
	"fmt"
	"log"
	"my_gin/models"
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
	//endless.DefaultReadTimeOut = setting.ReadTimeOut
	//endless.DefaultWriteTimeOut = setting.WriteTimeOut
	//endless.DefaultMaxHeaderBytes = 1 << 20 //请求头的最大字节数 2^20 即：1024 * 1024b => 1M
	//ipPort := fmt.Sprintf(":%d", setting.HttpPort)

	//server := endless.NewServer(ipPort, router)//windows系统没有 syscall.SIGUSR1 和 syscall.SIGUSR2，所以会报此错
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())//启动时输出pid
	//}
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", setting.ServerSetting.HttpIp,setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeOut,
		WriteTimeout:   setting.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20, //请求头的最大字节数 2^20 即：1024 * 1024b => 1M
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
