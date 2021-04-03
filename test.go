package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
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

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)

	//
	//var new VCard
	//json.Unmarshal(js,&new)
	//fmt.Println(*new.Addresses[1])
	//fmt.Println(*vc.Addresses[1])
	s := "eyJGaXJzdE5hbWUiOiJKYW4iLCJMYXN0TmFtZSI6IktlcnNzY2hvdCIsIkFkZHJlc3NlcyI6W3siVHlwZSI6InByaXZhdGUiLCJDaXR5IjoiQWFydHNlbGFhciIsIkNvdW50cnkiOiJCZWxnaXVtIn0seyJUeXBlIjoid29yayIsIkNpdHkiOiJCb29tIiwiQ291bnRyeSI6IkJlbGdpdW0ifV0sIlJlbWFyayI6Im5vbmUifQ=="
	//encoder := json.NewEncoder(os.Stdout)
	//encoder.Encode(js)
	//io.Reader()
	decoder := json.NewDecoder(strings.NewReader(s))
	decoder.Decode(js)
	fmt.Print(string(js))
}
