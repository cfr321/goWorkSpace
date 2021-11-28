package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

func (p Person) Pame() {
	fmt.Println("hello")
}

func (p Person) Nname() {
	fmt.Println("hello222")
}

func reflectTest(a interface{}, b interface{}) {
	v := reflect.ValueOf(a) // 或得 a 的值，然后可以查看具体信息和进行修改
	// v.Elem()     // It panics if v's Kind is not Interface or Ptr.  .Elem() = *ptr
	// v.Kind() //  类型标识

	for i := 0; i < v.NumMethod(); i++ { // 只能访问到大写的方法
		println(v.Type().Method(i).Name)
		v.Method(i).Call(nil)
	}
	for i := 0; i < v.Type().NumMethod(); i++ {
		println(v.Type().Method(i).Name)
	}

}
func other(a interface{}, b interface{}) {
	// 判断类型是否一样
	if reflect.TypeOf(a).Kind() == reflect.TypeOf(b).Kind() {

	}
	// 判断两个interface{}是否相等
	reflect.DeepEqual(a, b)

	// 将一个interface{}赋值给另一个interface{}
	reflect.ValueOf(a).Elem().Set(reflect.ValueOf(b))
}
func main() {
	m := make(map[int]string)
	var buf bytes.Buffer

	m[1] = "jjjj"
	m[23] = "4321432"

	// encode
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(&m)
	handErr(err)

	// decode
	decoder := gob.NewDecoder(&buf)
	err = decoder.Decode(&m)
	handErr(err)
	fmt.Println(m)
}

func handErr(err error) {
	if err != nil {
		fmt.Printf("%v", err)
	}
}
