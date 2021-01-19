package setting

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"my_gin/pkg/util"
)

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}

type RedisCfg struct {
	RedisInit   bool
	RedisClient map[string][]*redis.Client
	RedisList   map[string][]*RedisConfig `json:"redis"`
}
//加载store.json配置
func (this *RedisCfg) LoadRedisCfg(env string) error {
	file := util.NewFile()
	storeFile := fmt.Sprintf("conf/%s/store.json", env)
	jsonStr, _ := file.GetContentString(storeFile)
	err := json.Unmarshal([]byte(jsonStr), this)
	if err != nil {
		fmt.Printf("Could Unmarshal %s: %s\n", jsonStr, err)
		return err
	}

	if err = this.InitRedisList(); err != nil {
		fmt.Printf("InitRedisList error, err:%v!", err)
		return err
	}
	return nil
}

//初始化redis列表
func (this *RedisCfg) InitRedisList() error {
	if this.RedisInit { //已经初始化过了
		return nil
	}

	for redisName, redisData := range this.RedisList {
		var redisArr []*redis.Client
		for _, oneRedisCfg := range redisData {
			oneRedisCli := redis.NewClient(&redis.Options{
				Addr:     oneRedisCfg.Addr,
				Password: oneRedisCfg.Password,
				DB:       oneRedisCfg.Db,
				PoolSize: oneRedisCfg.PoolSize,
			})
			ret, _ := oneRedisCli.Ping().Result()
			if ret == "PONG" { //redis能连上
				redisArr = append(redisArr, oneRedisCli)
			} else {
				fmt.Printf("InitRedisList error, redis config:%v", redisName)
				return errors.New("InitRedisList error")
			}

		}
		//redis客户端列表
		this.RedisClient[redisName] = redisArr
	}
	this.RedisInit = true
	return nil
}

//获取redis列表中的redis客户端
func (this *RedisCfg)GetRedisClient(shareKey, name string) *redis.Client {
	if !this.RedisInit {
		if err := this.InitRedisList(); err != nil{
			return nil
		}
	}
	if _, ok := this.RedisClient[name]; ok {
		length := len(this.RedisList[name])
		if length > 0 {//有指定name的实例
			id := util.EncryptCRC32(shareKey)
			index := id % uint32(length)
			return this.RedisClient[name][index]
		}
	}
	return nil
}

