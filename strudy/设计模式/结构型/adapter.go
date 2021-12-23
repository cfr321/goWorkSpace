package main

import "fmt"

func compute(m mac) {
	m.sayhello()
}

type mac interface {
	sayhello()
}

type win struct {
}

func (w win) saywinhello() {
	fmt.Println("win say hello")
}

type adopter struct {
	w win
}

func (a adopter) sayhello() {
	a.w.saywinhello()
}

func main() {
	w := win{}
	a := adopter{w: w}
	compute(a)
}
