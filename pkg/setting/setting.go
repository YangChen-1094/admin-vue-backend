package setting

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type AppConfig struct {
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

type ServerConfig struct{
	RunMode string
	CorsUrl string
	HttpIp string
	HttpPort int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type DatabaseConfig struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
	MaxConn int
	MaxOpen int
}
type GrpcConfig struct {
	Ip string
	Port int
}
type Config struct {
	App	*AppConfig
	Server *ServerConfig
	Grpc *GrpcConfig
	Database *DatabaseConfig
}
var DeployConfig = &Config{
	App:&AppConfig{},
	Server: &ServerConfig{},
	Grpc: &GrpcConfig{},
	Database: &DatabaseConfig{},
}

//app.ini配置文件 变量
var Cfg *ini.File

//项目封装的加载配置 "Cfg"结尾的初始化对应配置
//“Config”结尾的配置的具体参数key-val
var Aws = &AwsCfg{}
var Redis = &RedisCfg{
	RedisClient: make(map[string][]*redis.Client),//如果不初始化 map（即通过make初始化），那么就会创建一个 nil map。nil map 不能用来存放键值对
}
var Mysql = &MysqlCfg{
	MysqlClient: make(map[string][]*gorm.DB),
}

func Setup(){
	var err error
	//initDeploy()
	runEnv := flag.String("env", "dev", "-env dev|pre|prod")//返回地址
	flag.Parse()
	appFile := fmt.Sprintf("conf/%s/app.ini", *runEnv)
	Cfg, err = ini.Load(appFile)
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	err = Cfg.Section("app").MapTo(DeployConfig.App)
	if err != nil {
		log.Fatalf("MapTo 'DeployConfig.App' Failed, err: %v", err)
	}
	DeployConfig.App.ImageMaxSize = DeployConfig.App.ImageMaxSize * 1024 * 1024//以M为单位

	err = Cfg.Section("server").MapTo(DeployConfig.Server)
	if err != nil {
		log.Fatalf("MapTo 'DeployConfig.Server' Failed, err: %v", err)
	}
	DeployConfig.Server.WriteTimeOut = DeployConfig.Server.WriteTimeOut * time.Second
	DeployConfig.Server.ReadTimeOut = DeployConfig.Server.ReadTimeOut * time.Second

	err = Cfg.Section("grpcConfig").MapTo(DeployConfig.Grpc)
	if err != nil {
		log.Fatalf("MapTo 'DeployConfig.Grpc' Failed, err: %v", err)
	}

	err = Cfg.Section("database").MapTo(DeployConfig.Database)
	if err != nil {
		log.Fatalf("MapTo 'DeployConfig.Database' Failed, err: %v", err)
	}

	//加载所有需要连接的db（包括redis、mysql、aws等）
	loadAllDb(*runEnv)
}

func GetExportPath() string{
	return fmt.Sprintf("%s", DeployConfig.App.ExportPath)
}

func initDeploy(){
	DeployConfig.App = &AppConfig{}
	DeployConfig.Server = &ServerConfig{}
	DeployConfig.Grpc = &GrpcConfig{}
	DeployConfig.Database = &DatabaseConfig{}
}

func loadAllDb(runEnv string){
	err := Redis.LoadRedisCfg(runEnv)
	if err != nil {
		log.Fatalf("loading 'Redis.LoadRedisCfg' Failed, err: %v", err)
	}
	err = Mysql.LoadMysqlCfg(runEnv)
	fmt.Println(Mysql)
	if err != nil {
		log.Fatalf("loading 'Redis.LoadMysqlCfg' Failed, err: %v", err)
	}

	err = Aws.LoadAwsCfg(runEnv)
	if err != nil {
		log.Fatalf("loading 'AwsConfig.LoadAwsCfg' Failed, err: %v", err)
	}
}