//
// Author: cfr
//

package main

import (
	"container/heap"
	"workspace/goLand/container"
)

type Int64 int64

func (a Int64) Cmp(cmpInterface container.CmpInterface) bool {
	return int64(a)> int64(cmpInterface.(Int64))
}

func main() {
	h := &container.Heap{Int64(1),Int64(2),Int64(3)}
	heap.Init(h)
	heap.Push(h,Int64(1))

}