package ttcp

import (
	"encoding/binary"
	fmt "fmt"
	"net"
)

type sessionMsg struct {
	number uint32
	length uint32
}

func TTcpServer(c chan struct{}) {
	server, err := net.Listen("tcp", ":8080")
	fmt.Println("stating ......")
	defer server.Close()
	handlerr(err)

	for true {
		conn, err := server.Accept()
		handlerr(err)
		var msg sessionMsg
		buf := make([]byte, 8)
		n, err := conn.Read(buf)
		handlerr2(err, n, 8)

		msg.number = binary.BigEndian.Uint32(buf[:4])
		msg.length = binary.BigEndian.Uint32(buf[4:8])

		buf = make([]byte, msg.length)
		var i uint32 = 0
		for ; i < msg.number; i++ {
			n, err := conn.Read(buf)
			handlerr(err)
			if n != int(msg.length) {
				_ = fmt.Errorf("read msg error no length")
			}
		}
		conn.Close()
	}
	c <- struct{}{}
}

func handlerr2(err error, n int, size int) {
	if err != nil || n != size {
		_ = fmt.Errorf("read msg error%v", err)
	}
}
func handlerr(err error) {
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
}
