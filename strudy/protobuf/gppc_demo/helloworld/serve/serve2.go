package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"math"
	"net"
	"workspace/strudy/protobuf/gppc_demo/helloworld/pb"
)

type server2 struct {
	Sycn     int
	w        int
	wordSize int
	times    []int
}

func (s *server2) SetTimeAndGetMsg(ctx context.Context, req *pb.Req) (*pb.Rep, error) {
	s.times[req.Rank] = int(req.Time)
	if s.Sycn == 0 {
		mi, mx := math.MaxInt32, 0
		for i := 0; i < s.wordSize; i++ {
			if s.times[i] > mx {
				mx = s.times[i]
			}
			if s.times[i] < mi {
				mi = s.times[i]
			}
		}
		if mx-mi > s.w {
			s.Sycn = s.wordSize
		}
	}
	syn := false
	if s.Sycn > 0 {
		syn = true
		s.Sycn--
		if s.Sycn == 0 {
			for i := 0; i < s.wordSize; i++ {
				s.times[i] = 0
			}
		}
	}
	return &pb.Rep{
		Syn: syn,
	}, nil
}

func newServer(wordSize int) *server2 {
	return &server2{
		Sycn:     0,
		w:        3,
		wordSize: wordSize,
		times:    make([]int, wordSize),
	}
}
func main() {

	conn, _ := net.Listen("tcp", ":9090")
	s := grpc.NewServer()
	pb.RegisterSynaCtrlServer(s, newServer(4))
	reflection.Register(s)
	_ = s.Serve(conn)
}
