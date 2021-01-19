package setting

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"my_gin/pkg/util"
	"time"
)

type MysqlConfig struct {
	Type string			`json:"type"`
	User string			`json:"user"`
	Password string		`json:"password"`
	Host string			`json:"host"`
	Name string			`json:"name"`
	TablePrefix string	`json:"table_prefix"`
	MaxConn int 		`json:"max_conn"`
	MaxOpen int			`json:"max_open"`
}
type MysqlCfg struct {
	MysqlInit   bool
	MysqlClient map[string][]*gorm.DB
	MysqlList   map[string][]*MysqlConfig `json:"mysql"`
}

//加载store.json配置
func (this *MysqlCfg) LoadMysqlCfg(env string) error {
	file := util.NewFile()
	storeFile := fmt.Sprintf("conf/%s/store.json", env)
	jsonStr, _ := file.GetContentString(storeFile)
	err := json.Unmarshal([]byte(jsonStr), this)
	if err != nil {
		fmt.Printf("Could Unmarshal %s: %s\n", jsonStr, err)
		return err
	}

	if err = this.InitMysqlList(); err != nil {
		fmt.Printf("InitRedisList error, err:%v!", err)
		return err
	}
	return nil
}

//初始化mysql list
func (this *MysqlCfg)InitMysqlList() error {
	if this.MysqlInit {
		return nil
	}
	for mysqlName, mysqlConf := range this.MysqlList {
		mysqlArr := make([]*gorm.DB, 0)
		for _, oneMysqlCfg := range mysqlConf {
			str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
				oneMysqlCfg.User,
				oneMysqlCfg.Password,
				oneMysqlCfg.Host,
				oneMysqlCfg.Name)
			mysqlDb, err := gorm.Open(oneMysqlCfg.Type, str) //这里必须使用 “=” 让db作为包内的变量 ，如果使用“:=”只能运用再Setup()方法内
			if err != nil {
				str := fmt.Sprintf("Fail to Open 'databaseList:%v ': %v", mysqlName, err)
				return errors.New(str)
			}
			mysqlDb.DB().SetMaxIdleConns(oneMysqlCfg.MaxConn)  //最大的空闲连接数
			mysqlDb.DB().SetMaxOpenConns(oneMysqlCfg.MaxOpen)  //最大的连接数
			mysqlArr = append(mysqlArr, mysqlDb)
		}
		this.MysqlClient[mysqlName] = mysqlArr
	}
	this.MysqlInit = true
	return nil
}

//获取redis列表中的redis客户端
func (this *MysqlCfg)GetMysqlClient(name, rangeId string) *gorm.DB {
	if !this.MysqlInit {
		if err := this.InitMysqlList(); err != nil{
			return nil
		}
	}
	if _, ok := this.MysqlClient[name]; ok {
		length := len(this.MysqlClient[name])
		if length > 0 {//有指定name的实例
			if length == 1{
				fmt.Println("GetMysqlClient index: 0")
				return this.MysqlClient[name][0]
			}
			var id uint32
			if rangeId == ""{
				rand.Seed(time.Now().UnixNano())//产生一个随机数种子，不然下面的方法都是返回同样的数字
				id = rand.Uint32()
			}else{
				id = util.EncryptCRC32(rangeId)
			}

			index := id % uint32(length)
			return this.MysqlClient[name][index]
		}
	}
	return nil
}

