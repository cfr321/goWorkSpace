package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strings"
	"time"
)

func rcp1() {

	net.Dial("tcp", "localhost:1234")

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Client dial fatal", err)
	}
	var reply string
	err = client.Call("HelloService.SayHello", "nihao", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply)
}

func echo(c net.Conn, shout string) {
	fmt.Fprintln(c, strings.ToUpper(shout))
	fmt.Fprintln(c, shout)
	fmt.Fprintln(c, strings.ToLower(shout))
	time.Sleep(time.Second)
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		echo(c, input.Text())
	}
	// 注意：忽略 input.Err() 中可能的错误
	i := []byte("byte")
	c.Write(i)
	c.Close()

}

//func handleConn0(c net.Conn) {
//	io.Copy(c, c)
//	c.Close()
//}

// RPC JSON
func main() {
	listen, _ := net.Listen("tcp", ":1234")
	for {
		conn, _ := listen.Accept()
		go handleConn(conn)
	}
}
