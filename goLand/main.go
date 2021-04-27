//
// Author: cfr
//

package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
	"workspace/goLand/container"
)

type Int64 int64

func (a Int64) Cmp(cmpInterface container.CmpInterface) bool {
	return int64(a) > int64(cmpInterface.(Int64))
}

func testReflect() {
	var a float64
	a = 1.1
	e := reflect.ValueOf(&a)
	e.Elem().SetFloat(2.3)
	fmt.Println(a)
}

func main() {
	//type T struct {
	//	A int
	//	B string
	//}
	//t := T{23, "skidoo"}
	//s := reflect.ValueOf(&t).Elem()
	//typeOfT := s.Type()
	//for i := 0; i < s.NumField(); i++ {
	//	f := s.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n", i,
	//		typeOfT.Field(i).Name, f.Type(), f.Interface())
	//}
	//var a sync.Map
	//now := time.Now()
	//now.Format("2006-01-02,15:04:05")
	//fmt.Println(now)
	//timer := time.NewTimer(5 * time.Second)
	//fmt.Println(time.Now())
	//<-timer.C
	//fmt.Println(time.Now())

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {

		log.Println("Ticker tick.")
	}
}
