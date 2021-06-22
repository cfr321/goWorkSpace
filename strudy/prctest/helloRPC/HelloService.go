package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type HelloService struct {
}

/*
 rpc: serve
		1、需要一个链接： HttpHand、 net.conn、
		2、需要注册方法： Register、RegisterName
		3、需要一种编码： json、普通
*/
func (s *HelloService) SayHello(req string, reply *string) error {
	*reply = req + "hello"
	return nil
}

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(out io.Writer, in io.Reader) {
	if _, err := io.Copy(out, in); err != nil {
		log.Fatal(err)
	}
}

func rpc1() {
	rpc.Register(new(HelloService))

	rpc.HandleHTTP()
	serve, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listenTCP fatal", err)
	}
	http.Serve(serve, nil)

}
