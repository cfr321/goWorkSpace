//
// Author: cfr
//

package container

import (
	"container/heap"
	"fmt"
)

type CmpInterface interface {
	Cmp(cmpInterface CmpInterface) bool
}

// IntHeap是一个整型的整数。
type Heap []CmpInterface
var one *Heap
func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Cmp(h[j])}
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push和Pop使用指针接收器，因为它们修改切片的长度，
	// 不仅仅是其内容。
	*h = append(*h, x.(CmpInterface))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewHeap() *Heap {
	*one = append((*one)[:0])
	return one
}


//使用实例
type Int int

func (a Int) Cmp(b CmpInterface) bool {
	return int(a) < int(b.(Int))
}
func main() {
	var b Heap
	heap.Init(&b)
	heap.Push(&b,Int(1))
	heap.Push(&b,Int(2))
	heap.Push(&b,Int(8))
	heap.Push(&b,Int(5))
	fmt.Print(heap.Pop(&b))
}
