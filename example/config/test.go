package main

import (
	"encoding/json"
	"fmt"
	"my_gin/pkg/file"
	"os"
)

type RedisCfg struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}

type AllDbConfig struct {
	RedisList map[string][]*RedisCfg	`json:"redis"`
	Etcd	  []string					`json:"etcd"`
}
var DbCfg AllDbConfig
func main(){
	dir,_ := os.Getwd()
	configFile := fmt.Sprintf("%s/test.json", "./conf/dev")
	contents, _ := file.GetContentString(dir + configFile)
	_ = json.Unmarshal([]byte(contents), &DbCfg)//把json字符串数据，解码到相应的数据结构


	fmt.Println(contents)

	fmt.Println(DbCfg)
}
