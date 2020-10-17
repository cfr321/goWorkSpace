//
// Author: cfr
//
package main
type ListNode struct {
	Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	var res,last *ListNode
	res=head
	l:=0
	for res!=nil {
		l++
		last=res
		res=res.Next
	}
	if l == 1 || l == 0 {
		return head
	}
	k%=l
	k=l-k
	if k == 0 {
		return head
	}
	res=head
	for i := 0; i < k-1; i++ {
		res=res.Next
	}
	last.Next=head
	head=res.Next
	res.Next=nil
	return head
}
