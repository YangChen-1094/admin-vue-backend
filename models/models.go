package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"my_gin/pkg/setting"
)

var db *gorm.DB

type Model struct{}

func Setup() {
	var err error
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	db, err = gorm.Open(setting.DatabaseSetting.Type, str) //这里必须使用 “=” 让db作为包内的变量 ，如果使用“:=”只能运用再Setup()方法内
	if err != nil {
		log.Fatalf("Fail to Open 'database': %v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true) //默认情况下使用单数表
	//db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDb() {
	defer db.Close()
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