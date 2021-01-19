package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"my_gin/pkg/setting"
)

var Db *gorm.DB

type Model struct{}

func Setup() {
	var err error
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DeployConfig.Database.User,
		setting.DeployConfig.Database.Password,
		setting.DeployConfig.Database.Host,
		setting.DeployConfig.Database.Name)
	Db, err = gorm.Open(setting.DeployConfig.Database.Type, str) //这里必须使用 “=” 让db作为包内的变量 ，如果使用“:=”只能运用再Setup()方法内
	if err != nil {
		log.Fatalf("Fail to Open 'database': %v", err)
	}
	gorm.DefaultTableNameHandler = func(Db *gorm.DB, defaultTableName string) string {
		return setting.DeployConfig.Database.TablePrefix + defaultTableName
	}

	Db.SingularTable(true) //默认情况下使用单数表
	//Db.LogMode(true)
	Db.DB().SetMaxIdleConns(setting.DeployConfig.Database.MaxConn)  //最大的空闲连接数
	Db.DB().SetMaxOpenConns(setting.DeployConfig.Database.MaxOpen)  //最大的连接数

	RedisMgr = NewRedisManager()
	RedisMgr.Init("127.0.0.1:6379")
}

func CloseDb() {
	defer Db.Close()
}

//model类
func NewModelTag() *ModelTag {
	return &ModelTag{}
}

func NewModelArticle() *ModelArticle {
	return &ModelArticle{}
}

func NewModelAuth() *ModelAuth{
	return &ModelAuth{}
}

func NewModelCaptcha() *ModelCaptcha{
	return &ModelCaptcha{}
}

func NewModelChannel() *ModelChannel{
	return &ModelChannel{}
}

func NewModelItem() *ModelItem{
	return &ModelItem{}
}
func NewModelCron()*ModelCron{
	return &ModelCron{}
}