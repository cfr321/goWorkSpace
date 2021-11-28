package mpc

import (
	"fmt"
	"os"
	"testing"
)

func TestInitWorkGroup(t *testing.T) {
	for i := 0; i < 4; i++ {
		go func(rank int) {
			err := InitWorkGroup(rank, 4)
			if err != nil {
				os.Exit(1)
			}
			var data []byte
			if rank == 0 {
				data = []byte("hello broadcast")
			}
			Broadcast(0, &data)
			fmt.Println(data)
		}(i)
	}
}
