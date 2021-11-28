package main

import (
	"fmt"
	"github.com/chenfar/syncserver"
)

func main() {
	server := syncserver.NewServer()
	fmt.Println("start")
	server.Start()
}
