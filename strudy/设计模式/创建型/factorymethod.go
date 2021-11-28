package main

import (
	"flag"
	"fmt"
)

type operater interface {
	SetA(a int)
	SetB(b int)
	Result() int
}

type baseOperater struct {
	a, b int
}

func (op *baseOperater) SetA(a int) {
	op.a = a
}

func (op *baseOperater) SetB(b int) {
	op.b = b
}

type sumOperater struct {
	baseOperater
}

func (op *sumOperater) Result() int {
	return op.a + op.b
}

type subOperater struct {
	baseOperater
}

func (op *subOperater) Result() int {
	return op.a - op.b
}

type opFactory interface {
	CreateOperater() operater
}
type sumOpFactory struct {
}

func (f *sumOpFactory) CreateOperater() operater {
	return &sumOperater{}
}

type subOpFactory struct {
}

func (f *subOpFactory) CreateOperater() operater {
	return &subOperater{}
}

func main() {
	s := flag.String("op", "sum", "")
	flag.Parse()
	var f opFactory
	switch *s {
	case "sub":
		f = &subOpFactory{}
	case "sum":
		f = &sumOpFactory{}
	default:
		panic("Not implement")
	}
	sumOp := f.CreateOperater()
	sumOp.SetA(1)
	sumOp.SetB(2)
	fmt.Println(sumOp.Result())
}
