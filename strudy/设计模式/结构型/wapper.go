package main

// 装饰器模式，装饰器持有一个接口对象，在调用接口方法时现调用自己持有的那个接口的对象的方法

type Food interface {
	getPrice() int
}

type Apple struct {
}

func (a *Apple) getPrice() int {
	return 5
}

type Wapper struct {
	food Food
}

func (w *Wapper) getPrice() int {
	p := w.food.getPrice()
	return p + 5
}
