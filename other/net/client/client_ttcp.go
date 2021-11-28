package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	number int
	length int
)

func init() {
	flag.IntVar(&number, "number", 100, "the number to send")
	flag.IntVar(&length, "length", 1024, "the length to send")
}

func main() {
	response, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(resp))
}

func client() {
	//net.ResolveTCPAddr()

	flag.Parse()
	fmt.Println(number, length)
	conn, err := net.Dial("tcp", "localhost:8080")
	handlerr(err)
	fmt.Println("connect success")
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint32(buf[:4], uint32(number))
	binary.BigEndian.PutUint32(buf[4:], uint32(length))
	n, err := conn.Write(buf[:8])
	handlerr(err)
	if n != 8 {
		fmt.Println("not e")
	}
	buf = make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = "0123456789ABCDEF"[i%16]
	}
	fmt.Println(number, length)

	begin := time.Now()
	for i := 0; i < number; i++ {
		n, err := conn.Write(buf)
		if err != nil || n != length {
			_ = fmt.Errorf("%v", err)
			return
		}
	}
	usetime := time.Since(begin).Seconds()
	fmt.Printf("using time: %v\n", usetime)
	fmt.Println("speed:", 1.0*float64(number)*float64(length)/1024/1024/usetime, "MB/S")
	conn.Close()
}
func handlerr(err error) {
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
}
