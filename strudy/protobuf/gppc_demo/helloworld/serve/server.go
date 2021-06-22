package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
	"workspace/strudy/protobuf/gppc_demo/helloworld/pb"
)

type server struct {

}
var time int
func (s *server)SayHello(cx context.Context,request *pb.HelloRequest) (*pb.HelloReply,error) {
	time ++
	return &pb.HelloReply{
		Message: "Hello" + request.Name + strconv.Itoa(time),
	},nil
}
func main() {
	time = 0
	lis, _ := net.Listen("tcp", ":9090")
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})
	reflection.Register(s)
	err := s.Serve(lis)
	if err != nil {
		return
	}
}
