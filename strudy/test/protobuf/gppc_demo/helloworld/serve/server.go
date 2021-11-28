package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
	"workspace/strudy/test/protobuf/gppc_demo/helloworld/pb"
)

//$ protoc --go_out=. ./helloworld.proto
//$ protoc --go_out=plugins=grpc:.  ./helloworld.proto
/**
边两个命令是不一样的。第一个只是生成了proto序列化，反序列化代码的文件。
第二个则还增加服务器和客户端通讯、实现的公共库代码。因此如果是写客户端和服务端通信，
要用第二个编译方式，如果只是作为序列化和反序列化的工具，用第一个就可以了。
*/

type server struct {
}

var time int

func (s *server) SayHello(cx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	time++
	return &pb.HelloReply{
		Message: "Hello" + request.Name + strconv.Itoa(time),
	}, nil
}
func main() {
	time = 0
	lis, _ := net.Listen("tcp", ":9090")
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	err := s.Serve(lis)
	if err != nil {
		return
	}
}
