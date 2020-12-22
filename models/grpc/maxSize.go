package modelGrpc

import (
	"context"
	"log"
	"my_gin/proto"
)

type MaxSize struct {
}

func (this *MaxSize) One(context.Context, *proto.Empty) (*proto.StringSingle, error){
	log.Println("[server grpc] [one]")
	return &proto.StringSingle{}, nil
}
func (this *MaxSize) ClientStream(stream proto.MaxSize_ClientStreamServer) error{
	return nil
}
func (this *MaxSize) ServerStream(*proto.Empty, proto.MaxSize_ServerStreamServer) error{
	return nil
}

func (this *MaxSize) DoubleStream(stream proto.MaxSize_DoubleStreamServer) error{
	return nil
}
