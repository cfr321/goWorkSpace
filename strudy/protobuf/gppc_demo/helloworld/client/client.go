package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"math/rand"
	"sync"
	"time"
	"workspace/strudy/protobuf/gppc_demo/helloworld/pb"
)

var wg sync.WaitGroup

func main() {

	//client := pb.NewSynaCtrlClient(con)
	//wg = sync.WaitGroup{}
	//for i := 0; i < 4; i++ {
	//	go todo(i, client)
	//	wg.Add(1)
	//}
	//wg.Wait()

	for i := 0; i < 2; i++ {
		go callSayHello(i)
	}
	c := make(chan struct{})
	<-c
}

func todo(rank int, client pb.SynaCtrlClient) {
	syncTime := 0
	epoch := 0
	for {
		epoch++
		msg, err := client.SetTimeAndGetMsg(context.Background(), &pb.Req{Time: int32(epoch), Rank: int32(rank)})
		if err != nil {
			fmt.Println(err)
			wg.Done()
		}
		if msg.Syn {
			syncTime++
			fmt.Printf("sync at rank %d, in epoch %d, in syncTime %d\n", rank, epoch, syncTime)
			epoch = 0
			time.Sleep(5 * time.Second)
		}
		n := rand.Int31n(2)
		time.Sleep(time.Duration(n) * time.Second)
	}
	wg.Done()
}

func callSayHello(rank int) {
	con, _ := grpc.Dial(":9090", grpc.WithInsecure())
	defer con.Close()
	client := pb.NewGreeterClient(con)
	for i := 0; i < 10000; i++ {
		_, _ = client.SayHello(context.Background(), &pb.HelloRequest{
			Name: "name",
		})
	}
}
