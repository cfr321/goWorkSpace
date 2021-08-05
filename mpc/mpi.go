package mpc

import (
	"github.com/cfr321/mpc/comm"
)

type MpiError string

func (v MpiError) Error() string {
	return "MpiError: " + string(v)
}

func InitWorkGroup(rank int, worldSize int) error {
	return initWorkGroup(rank, worldSize)
}

func Barrier() {
	barrier(getrank())
}

func Send(to int, data []byte) {
	send(to, data)
}

func Recv() []byte {
	return recv(getrank())
}

func AllReduce(data []byte) [][]byte {
	return allreduce(getrank(), data)
}

func AllReduceFloat(data float64) float64 {
	reduce := AllReduce(comm.Float64ToByte(data))
	res := 0.0
	for i := 0; i < len(reduce); i++ {
		res += comm.ByteToFloat64(reduce[i])
	}
	return res
}

func Broadcast(src int, data *[]byte) {
	rank := getrank()
	broadcast(rank, src, data)
}

func Scatter(src int, data *[]byte) {
	self := getrank()
	scatter(self, src, data)
}

func Gather() {

}

func WorldSize() int {
	return worldsize()
}

func Rank() int {
	return getrank()
}
