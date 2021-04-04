//
// Author: cfr
//

package myUtil

type CmpInterface interface {
	Less(other CmpInterface) bool
}
type priorityQueue struct {
	items []CmpInterface
}

func (h *priorityQueue) Len() int {
	return len(h.items)
}
func (h *priorityQueue) Less(i, j int) bool {
	return h.items[i].Less(h.items[j])
}
func (h *priorityQueue) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
func (h *priorityQueue) Push(x interface{}) {
	h.items = append(h.items, x.(CmpInterface))
}
func (h *priorityQueue) Pop() interface{} {
	if len(h.items) == 0 {
		return nil
	}
	res := h.items[len(h.items)-1]
	h.items = append(h.items[:len(h.items)-1])
	return res
}
func NewHeap() *priorityQueue {
	return &priorityQueue{
		items: []CmpInterface{},
	}
}
