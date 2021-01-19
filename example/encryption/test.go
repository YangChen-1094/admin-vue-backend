package main

import (
	"fmt"
	"my_gin/pkg/setting"
	"time"
)

func main(){
	setting.Setup()
	redisGame := setting.Redis.GetRedisClient("test", "game")
	thatTime, _:= time.ParseInLocation("2006-01-02 15:04:05", "2021-05-11 14:06:06", time.Local)
	expire, _ := time.ParseDuration("30h")//30小时之后
	fmt.Println("expire:", expire)
	fmt.Println("ret:", thatTime.Unix())
	redisGame.ExpireAt("test", thatTime)
}