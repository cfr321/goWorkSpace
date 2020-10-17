//
// Author: cfr
//

package myUtil

type intHeap struct {
	items []int
	L     bool
}
func (h intHeap) Len() int {
	return len(h.items)
}
func (h intHeap) Less(i, j int) bool {
	if h.L {
		return h.items[i] < h.items[j]
	} else {
		return h.items[i] > h.items[j]
	}
}
func (h intHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
func (h *intHeap) Push(x interface{}) {
	h.items = append(h.items, x.(int))
}
func (h *intHeap) Pop() interface{} {
	res := h.items[len(h.items)-1]
	h.items = append(h.items[:len(h.items)-1])
	return res
}
func NewIntHeap(Less bool) *intHeap {
	return &intHeap{
		items: []int{},
		L:     Less,
	}
}
