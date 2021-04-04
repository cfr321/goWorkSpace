package class

import "fmt"

type Person struct {
	name string
}

func (p Person) printName() {
	fmt.Print(p.name)
}

type Student struct {
	Person
	sex string
}

//func main() {
//	var s Student
//	s.name = "xiaoming"
//	s.printName()
//}
