package setting

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret string
	PageSize int
	RuntimeRootPath string
	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize	int
	ImageAllowExts	[]string
	ExportPath string
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	LogTimeFormat string
}

type Server struct{
	RunMode string
	CorsUrl string
	HttpIp string
	HttpPort int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
	MaxConn int
	MaxOpen int
}
type Grpc struct {
	Ip string
	Port int
}

var AppSetting = &App{}
var ServerSetting = &Server{}
var GrpcSetting = &Grpc{}
var DatabaseSetting = &Database{}
var Cfg *ini.File

func Setup(){
	var err error
	runEnv := flag.String("env", "dev", "-env dev|pre|prod")//返回地址
	appFile := fmt.Sprintf("conf/%s/app.ini", *runEnv)
	Cfg, err = ini.Load(appFile)
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("MapTo 'AppSetting' Failed, err: %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024//以M为单位

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("MapTo 'ServerSetting' Failed, err: %v", err)
	}
	ServerSetting.WriteTimeOut = ServerSetting.WriteTimeOut * time.Second
	ServerSetting.ReadTimeOut = ServerSetting.ReadTimeOut * time.Second

	err = Cfg.Section("grpcConfig").MapTo(GrpcSetting)
	if err != nil {
		log.Fatalf("MapTo 'GrpcSetting' Failed, err: %v", err)
	}

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("MapTo 'DatabaseSetting' Failed, err: %v", err)
	}
}