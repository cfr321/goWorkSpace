package main

import (
	"fmt"
	"strconv"
)

func main() {
	res :=  strconv.FormatInt(5, 2)
	fmt.Print(fmt.Sprintf("%032s",res))
}