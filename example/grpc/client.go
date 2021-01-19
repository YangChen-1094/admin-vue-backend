package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"my_gin/models"
	"my_gin/pkg/setting"
	"my_gin/proto"
)

func init() {//初始化
	setting.Setup()
	models.Setup()
}

func main(){
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", setting.DeployConfig.Grpc.Ip, setting.DeployConfig.Grpc.Port), grpc.WithInsecure())
	log.SetFlags(log.Llongfile | log.LstdFlags | log.Lmicroseconds)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewMaxSizeClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	one(ctx, client)
}

func one(ctx context.Context, client proto.MaxSizeClient) {
	///一元RPC
	stream, err := client.One(ctx, &proto.Empty{Name: "this is client One"})
	log.Println("[client] one():", stream, err)
	log.Println()
	log.Println()
}
