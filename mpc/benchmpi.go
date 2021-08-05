package mpc

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type group struct {
	worldSize int
	recvChans []chan []byte
	locker    sync.Mutex
	rankMap   map[string]int
}

type block struct {
	blockChans []chan struct{}
	num        int32
}

var workGroup *group
var blk *block
var once sync.Once

func initWorkGroup(rank int, worldSize int) error {
	if rank < 0 || worldSize < 0 || rank >= worldSize {
		return MpiError("worldSize must > 0, 0 <= rank < worldSize")
	}
	once.Do(func() {
		workGroup = &group{
			worldSize: worldSize,
			recvChans: make([]chan []byte, worldSize),
			locker:    sync.Mutex{},
			rankMap:   make(map[string]int),
		}
		blk = &block{
			blockChans: make([]chan struct{}, worldSize),
			num:        0,
		}
	})
	if worldSize != workGroup.worldSize {
		return MpiError("worldSize must same")
	}
	workGroup.locker.Lock()
	if workGroup.recvChans[rank] != nil {
		return MpiError(fmt.Sprintf("rank %d had inited", rank))
	}
	workGroup.recvChans[rank] = make(chan []byte, worldSize*2)
	workGroup.rankMap[groutineName()] = rank
	workGroup.locker.Unlock()
	blk.blockChans[rank] = make(chan struct{}, 1)

	barrier(rank)

	return nil
}

func barrier(self int) {
	atomic.AddInt32(&blk.num, 1)
	if atomic.CompareAndSwapInt32(&blk.num, int32(worldsize()), 0) {
		for i := 0; i < worldsize(); i++ {
			blk.blockChans[i] <- struct{}{}
		}
	}
	<-blk.blockChans[self]
}

func broadcast(self int, src int, data *[]byte) {
	if self == src {
		for i := 0; i < len(workGroup.recvChans); i++ {
			workGroup.recvChans[i] <- *data
		}
	}
	tmp := <-workGroup.recvChans[self]
	*data = make([]byte, len(tmp))
	copy(*data, tmp)
}

func send(to int, data []byte) {
	workGroup.recvChans[to] <- copybyte(data)
}

func recv(self int) []byte {
	return <-workGroup.recvChans[self]
}

func scatter(self int, src int, data *[]byte) {
	//sliceSize := len(*data) /
}

func gather(self int, to int, data *[]byte) {

}

func allreduce(self int, data []byte) [][]byte {
	ws := worldsize()
	next := (self + 1) % ws
	data = copybyte(data)
	var res [][]byte
	res = append(res, data)
	for i := 0; i < ws-1; i++ {
		send(next, data)
		data = recv(self)
		res = append(res, data)
	}
	return res
}

func copybyte(data []byte) []byte {
	tmp := make([]byte, len(data))
	copy(tmp, data)
	return tmp
}

func getrank() int {
	b := groutineName()
	return workGroup.rankMap[b]
}

func worldsize() int {
	return workGroup.worldSize
}

func groutineName() string {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	return string(b)
}
