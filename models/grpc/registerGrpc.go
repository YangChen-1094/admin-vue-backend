package modelGrpc

import (
	"fmt"
	"google.golang.org/grpc"
	"math"
	"my_gin/pkg/setting"
	"my_gin/proto"
	"net"
)

//注册GRPC服务
func Register(){
	//注册pb的协议
	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(math.MaxInt32),
		grpc.MaxSendMsgSize(math.MaxInt32),
	}
	grpcServer := grpc.NewServer(options...)
	grpcMyModel := &MaxSize{}//server端实现pb协议中MaxSize grpc的方法
	proto.RegisterMaxSizeServer(grpcServer, grpcMyModel)

	go func() {//必须以协程方式启动，不然gin框架启动不了
		grpcCon, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", setting.GrpcSetting.Ip, setting.GrpcSetting.Port))
		_ = grpcServer.Serve(grpcCon)
		defer grpcCon.Close()
	}()
}
