package main

import (
	"fmt"
	"my_gin/pkg/util"
)

func main(){
	before := "YangChen123"
	num := util.EncryptSHA1(before)
	fmt.Println("sha1 前：", before,"加密：", num)
}