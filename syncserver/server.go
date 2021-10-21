package syncserver

import (
	"context"
	"github.com/chenfar/syncserver/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync/atomic"
)

const (
	DefaultWorldSize = 4
	MAXE             = 8
)

type Server struct {
	worldSize uint32
	finished  []bool
	overNum   uint32
	maxE      uint32
	callNum   uint32
}

func NewServer() *Server {
	return &Server{
		worldSize: DefaultWorldSize,
		finished:  make([]bool, DefaultWorldSize),
		overNum:   0,
		maxE:      MAXE,
		callNum:   0,
	}
}

func (s *Server) SyncCtl(ctx context.Context, req *protoc.Res) (*protoc.Resp, error) {
	rank := req.GetRank()
	epoch := req.GetEpoch()
	if !s.finished[rank] {
		s.finished[rank] = true
		atomic.AddUint32(&s.overNum, 1)
	}
	var res *protoc.Resp
	if s.overNum == s.worldSize || epoch == s.maxE {
		res = &protoc.Resp{
			Sync: true,
		}
	} else {
		res = &protoc.Resp{
			Sync: false,
		}
	}
	if res.Sync == true {
		atomic.AddUint32(&s.callNum, 1)
	}
	if s.callNum == s.worldSize {
		go s.reset()
	}
	return res, nil
}

func (s *Server) Start() {
	listen, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalln("start server", err)
	}
	server := grpc.NewServer()
	protoc.RegisterSyncServerServer(server, s)
	reflection.Register(server)
	err = server.Serve(listen)
	if err != nil {
		log.Fatalln("start server", err)
	}
}

func (s *Server) reset() {
	s.finished = make([]bool, DefaultWorldSize)
	s.callNum = 0
	s.overNum = 0
}
