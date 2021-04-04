package main

import (
	"fmt"
	"math"
	"net/http"
	"workspace/myUtil"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func trap(height []int) int {
	var s []int
	res := 0
	//s = append(s, 0)
	for i := 0; i < len(height); i++ {
		for len(s) > 0 && height[i] >= height[s[len(s)-1]] {
			t := s[len(s)-1]
			s = s[0 : len(s)-1]
			if len(s) > 0 {
				before := s[len(s)-1]
				res += (min(height[i], height[before]) - height[t]) * (i - before)
			}
		}
		s = append(s, i)
	}
	return res
}

func trap1(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0

	for left < right {
		leftMax = int(math.Max(float64(leftMax), float64(height[left])))
		rightMax = int(math.Max(float64(rightMax), float64(height[right])))
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

type Node struct {
	a, b int
}

func (n Node) Less(other myUtil.CmpInterface) bool {
	return n.b < other.(Node).b
}

type easy struct {
	Name string `json:"name"` // 字段解释，可指json 字符串的名字
	Age  int    `json: age`
	Like string `json: like`
}

func handle(writer http.ResponseWriter, request *http.Request) {
	data := []byte("helloNihao")
	writer.Write(data)
}

func handle2(writer http.ResponseWriter, request *http.Request) {
	data := []byte("草拟吗")
	writer.Write(data)
}

func main() {
	http.HandleFunc("/go", handle)
	http.HandleFunc("/hello", handle2)
	fmt.Print("run app")
	http.ListenAndServe(":1234", nil)
}
