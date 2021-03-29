//
// Author: cfr
//
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	p := head
	len := 1
	for p.Next != nil {
		len++
		p = p.Next
	}
	if len == 0 && len == 1 {
		return head
	}

	k %= len
	k = len - k
	res := head
	for k > 1 {
		res = res.Next
		k--
	}
	p.Next = head
	head = res.Next
	res.Next = nil
	return head
}
func addSlice(arr *[]int) {
	*arr = append(*arr, 1)
}
func main() {

	ints := make([]int, 10)
	addSlice(&ints)
	fmt.Print(ints)
}
