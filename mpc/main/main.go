package main

import (
	"fmt"
	"github.com/chenfar/mpc"

	"os"
	"sync"
)

func main() {
	group := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		group.Add(1)
		go func(rank int) {
			err := mpc.InitWorkGroup(rank, 4)
			fmt.Println("init finish ", rank)
			if err != nil {
				os.Exit(1)
			}
			mpc.Barrier()
			data := 100.0
			reduce := mpc.AllReduceFloat(data)
			fmt.Println(reduce)
			group.Done()
		}(i)
	}
	group.Wait()
}
