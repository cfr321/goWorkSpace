package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"time"
	"workspace/other/net/comm"
)

type messega struct {
	begin uint64
	resp  uint64
}

func (m *messega) encode() []byte {
	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.BigEndian, m.begin)
	comm.Handerr(err)
	err = binary.Write(buf, binary.BigEndian, m.resp)
	comm.Handerr(err)
	return buf.Bytes()
}

func (m *messega) decode(b []byte) {
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.BigEndian, &m.begin)
	comm.Handerr(err)
	err = binary.Read(buf, binary.BigEndian, &m.resp)
	comm.Handerr(err)
}
func now() uint64 {
	return uint64(time.Now().UnixMicro())
}
func main() {
	var c bool
	flag.BoolVar(&c, "c", false, "run as client,default server")
	flag.Parse()
	if c {
		client()
	} else {
		server()
	}
}

func client() {
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	defer conn.Close()
	comm.Handerr(err)
	go func() {
		for true {
			m := messega{0, 0}
			m.begin = now()
			nw, err := conn.Write(m.encode())
			//fmt.Println("send data")
			comm.Handerr(err)
			if nw != 16 {
				fmt.Println("err send ", nw)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	fmt.Println("get data....")
	size := binary.Size(messega{})
	buf := make([]byte, size)
	for true {
		var rem messega
		rn, err := conn.Read(buf)
		comm.Handerr(err)
		if rn < size {
			fmt.Println("err read ", rn)
			continue
		}
		rem.decode(buf)
		mytime := (now() + rem.begin) / 2
		fmt.Printf("my time :%v, server time: %v\n", mytime, rem.resp)
	}
}

func server() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	comm.Handerr(err)
	listenUDP, err := net.ListenUDP("udp", udpAddr)
	defer func() {
		listenUDP.Close()
		fmt.Println("main over!")
	}()
	comm.Handerr(err)
	size := binary.Size(messega{})
	buf := make([]byte, size)
	fmt.Println("begin to listen....")
	for true {
		var msg messega
		revsize, addr, err := listenUDP.ReadFrom(buf)
		//fmt.Println("get a data....")
		comm.Handerr(err)
		if revsize == size {
			msg.decode(buf)
			msg.resp = now()
			data := msg.encode()
			to, err := listenUDP.WriteTo(data, addr)
			comm.Handerr(err)
			if to < 0 {
				_ = fmt.Errorf("send to %v error\n", addr)
			}
		} else {
			_ = fmt.Errorf("read from %v error\n", addr)
		}
	}
}
