# container容器包里面三个容器的使用

## heap
go语言没有实现开箱即用的堆，但实现了堆的接口，需要你实现接口。

下面是一个我基于heap的一个封装。
```go
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
```

原理也很简答，我重新写了一个接口，但是这个接口只要你实现cmp方法。那么任何实现了CmpInterface接口的类型，都可以放到堆里面来，形成一个优先队列。

## List是一个双向链表

这是一个开箱可用的容易，简单了解它的使用规则就好了。      

在语言的自带链表，通常都设计成双向链表。主要的函数有以下这些

type Element

- func (e *Element) Next() *Element 

- func (e *Element) Prev() *Element 

type List

- func New() *List 

- func (l *List) Back() *Element 

- func (l *List) Front() *Element 

- func (l *List) Init() *List 

- func (l *List) InsertAfter(v interface{}, mark *Element) *Element 

- func (l *List) InsertBefore(v interface{}, mark *Element) *Element 

- func (l *List) Len() int 

- func (l *List) MoveAfter(e, mark *Element) 

- func (l *List) MoveBefore(e, mark *Element) 

- func (l *List) MoveToBack(e *Element) 

- func (l *List) MoveToFront(e *Element) 

- func (l *List) PushBack(v interface{}) *Element 

- func (l *List) PushBackList(other *List) 

- func (l *List) PushFront(v interface{}) *Element 

- func (l *List) PushFrontList(other *List) 

- func (l *List) Remove(e *Element) interface{} 



## Ring环

这是一个环形两双向链表

