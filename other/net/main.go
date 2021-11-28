package main

import (
	"fmt"
	"workspace/other/net/ttcp"
)

func main() {
	c := make(chan struct{})
	go ttcp.TTcpServer(c)
	<-c
	fmt.Println("finish")
}
