package main

import (
	"sort"
	"strconv"
	"strings"
)

// 01.01
func isUnique(astr string) bool {
	bytes := []byte(astr)
	sort.Slice(bytes, func(i, j int) bool {
		b := bytes[i] < bytes[j]
		return b
	})
	for i := 0; i < len(bytes)-1; i++ {
		if bytes[i] == bytes[i+1] {
			return false
		}
	}
	return true
}

func isUnique2(astr string) bool {
	var rem int
	for i := 0; i < len(astr); i++ {
		pos := astr[i] - 'a'
		tmp := 1 << pos
		tmp |= rem
		if tmp != 0 {
			return false
		}
		rem |= tmp
	}
	return true
}

// 01.03
func replaceSpaces(S string, length int) string {
	bytes := []byte(S)
	pos := len(S) - 1
	for i := length - 1; i >= 0; i-- {
		if bytes[i] == ' ' {
			bytes[pos] = '0'
			pos--
			bytes[pos] = '2'
			pos--
			bytes[pos] = '%'
			pos--
		} else {
			bytes[pos] = bytes[i]
			pos--
		}
	}
	return string(bytes[pos:])
}

// 01.04
func canPermutePalindrome(s string) bool {
	rem := make([]int, 128)
	for i := 0; i < len(s); i++ {
		rem[s[i]]++
	}
	res := 0
	for i := 0; i < 128; i++ {
		if rem[i]&1 != 0 {
			res++
		}
	}
	return res < 2
}

// 01.05
func oneEditAway(first string, second string) bool {
	lena := len(first)
	lenb := len(second)
	if lena-lenb > 1 || lenb-lena > 1 {
		return false
	}
	a, b := 0, 0
	flag := 0
	for a < lena && b < lenb {
		if first[a] == second[b] {
			a++
			b++
		} else {
			if flag == 1 {
				return false
			}
			flag++
			if lena == lenb {
				a++
				b++
			} else if lena > lenb {
				a++
			} else {
				b++
			}
		}
	}
	return true

	//if math.Abs(float64(len(first)-len(second))) >= 2 {
	//	return false
	//}
	//
	//dp := make([][]int, len(first)+1)
	//for i := 0; i < len(dp); i++ {
	//	dp[i] = make([]int,len(second)+1)
	//}
	//for i := 0; i <= len(first); i++ {
	//	dp[i][0] = i
	//}
	//for i := 0; i < len(second); i++ {
	//	dp[0][i] = i
	//}
	//for i := 0; i < len(first); i++ {
	//	for j := 0; j < len(second); j++ {
	//		if first[i] == second[j] {
	//			dp[i+1][j+1] = dp[i][j]
	//		}else{
	//			dp[i+1][j+1] = min(min(dp[i][j],dp[i+1][j]),dp[i][j+1])+1
	//		}
	//	}
	//}
	//return dp[len(first)][len(second)] < 2
}

//01.06
func compressString(S string) string {
	var res []byte
	var l = 1
	for i := 0; i < len(S); i++ {
		if i == len(S)-1 || S[i] != S[i+1] {
			res = append(res, S[i])
			res = append(res, strconv.Itoa(l)...)
			l = 1
		} else {
			l++
		}
	}
	if len(res) >= len(S) {
		return S
	}
	return string(res)
}

//01.07
func swap1(a, b *int) {
	*a, *b = *b, *a
}
func rotate(matrix [][]int) {
	N := len(matrix)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			swap1(&matrix[i][j], &matrix[j][i])
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N/2; j++ {
			swap1(&matrix[i][j], &matrix[i][N-j-1])
		}
	}
}

//01.08
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

//01.09   !!!!!!
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	new := s2 + s2
	index := strings.Index(new, s1)
	return index != -1

	//for i := 0; i < len(s1); i++ {
	//	tmp:=s1[i:]+s1[0:i]
	//	if tmp == s2 {
	//		return true
	//	}
	//}
	//return false
}
func test() {
	// add something to test
}

// 02.01
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var rem [200001]bool
	rem[head.Val] = true
	pre := head
	for pre.Next != nil {
		if rem[pre.Next.Val] {
			pre.Next = pre.Next.Next
		} else {
			rem[pre.Next.Val] = true
			pre = pre.Next
		}
	}
	return head
}

// 02.02
func kthToLast(head *ListNode, k int) int {
	var front, back = head, head
	for front != nil {
		front = front.Next
		k--
		if k < 0 {
			back = back.Next
		}
	}
	return back.Val
}

// 02.03
func deleteNode(node *ListNode) {
	if node.Next == nil {
		node = nil
	} else {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}

// 02.08
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast == head || slow != fast {
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		slow = slow.Next
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// 03.02
type StackOfPlates struct {
	cap    int
	stacks [][]int
}

func Constructor5(cap int) StackOfPlates {
	return StackOfPlates{cap: cap, stacks: [][]int{}}
}

func (this *StackOfPlates) Push(val int) {
	l := len(this.stacks)
	if l == 0 || len(this.stacks[l-1]) == this.cap {
		this.stacks = append(this.stacks, []int{})
	}
	l = len(this.stacks)
	this.stacks[l-1] = append(this.stacks[l-1], val)
}

func (this *StackOfPlates) Pop() int {
	l := len(this.stacks)
	if l == 0 {
		return -1
	}
	l1 := len(this.stacks[l-1])
	re := this.stacks[l-1][l1-1]
	if l1 == 1 {
		this.stacks = this.stacks[0 : l-1]
	} else {
		this.stacks[l-1] = this.stacks[l-1][0 : l1-1]
	}
	return re
}

func (this *StackOfPlates) PopAt(index int) int {
	l := len(this.stacks)
	if l == 0 || index > l-1 {
		return -1
	}
	l1 := len(this.stacks[index])
	re := this.stacks[index][l1-1]
	if l1 == 1 {
		this.stacks = append(this.stacks[0:index], this.stacks[index+1:l]...)
	} else {
		this.stacks[index] = this.stacks[index][0 : l1-1]
	}
	return re
}
