package main

import (
	"fmt"
)

type test struct {
	a int
}

func (t *test) testPrit() {
	t.a = 34
}
func main() {
	a := test{55}
	a.testPrit()
	fmt.Print(a.a)
}
