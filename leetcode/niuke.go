package main

import (
	"fmt"
	"sort"
)

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	p := pHead
	pHead = nil
	var tmp *ListNode
	for p != nil {
		tmp = p.Next
		p.Next = pHead
		pHead = p
		p = tmp
	}
	return pHead
}

type (
	LruNode struct {
		Key  int
		Val  int
		Pre  *LruNode
		Next *LruNode
	}
	LruMap struct {
		k    int
		m    map[int]*LruNode
		head *LruNode
		tail *LruNode
	}
)

func NewLruMap(k int) *LruMap {
	tmp := &LruMap{
		k:    k,
		m:    make(map[int]*LruNode),
		head: new(LruNode),
		tail: new(LruNode),
	}
	tmp.head.Next = tmp.tail
	tmp.tail.Pre = tmp.head
	return tmp
}

func (lru *LruMap) get(key int) int {
	if lruNode, has := lru.m[key]; has {
		lru.delete(lruNode)
		lru.addHead(lruNode)
		return lruNode.Val
	}
	return -1
}
func (lru *LruMap) put(key, value int) {
	if lruNode, has := lru.m[key]; has {
		lru.delete(lruNode)
		lruNode.Val = value
		lru.addHead(lruNode)
	} else {
		newNode := &LruNode{
			Key: key,
			Val: value,
		}
		lru.addHead(newNode)
		lru.m[key] = newNode

		// 如果超过了 k 个删除掉最后的
		if len(lru.m) > lru.k {
			deleteNode := lru.tail.Pre
			lru.delete(deleteNode)
			delete(lru.m, deleteNode.Key)
		}
	}
}

func (lru *LruMap) delete(node *LruNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (lru *LruMap) addHead(node *LruNode) {
	node.Next = lru.head.Next
	node.Next.Pre = node
	node.Pre = lru.head
	lru.head.Next = node
}

func LRU(operators [][]int, k int) (res []int) {
	// write code here
	lruMap := NewLruMap(k)
	for _, operator := range operators {
		if operator[0] == 1 {
			lruMap.put(operator[1], operator[2])
		} else {
			i := lruMap.get(operator[1])
			res = append(res, i)
		}
	}
	return
}

func hasCycle(head *ListNode) bool {
	// write code here
	fast := head
	slow := head

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
	}
	return false
}

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		for len(stack1) > 0 {
			stack2 = append(stack2, stack1[len(stack1)-1])
			stack1 = stack1[:len(stack1)-1]
		}
	}
	res := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return res
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// write code here
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func levelOrder(root *TreeNode) (res [][]int) {
	// write code here
	var queue []*TreeNode
	last := root
	queue = append(queue, root)
	var tmp []int
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		tmp = append(tmp, root.Val)
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if root == last {
			res = append(res, tmp)
			tmp = make([]int, 0)
			if len(queue) > 0 {
				last = queue[len(queue)-1]
			}
		}
	}
	return
}
func maxLength1(arr []int) int {
	rem := make(map[int]int)
	// write code here
	res := 0
	begin := 0
	for i := 0; i < len(arr); i++ {
		if index, has := rem[arr[i]]; has {
			for ; begin <= index; begin++ {
				delete(rem, arr[begin])
			}
		} else {
			rem[arr[i]] = i
		}
		res = max(res, i-begin+1)
	}
	return res
}

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here
	if root == nil {
		return -1
	}
	var dfs func(treeNode *TreeNode, f1 *int, f2 *int) int
	findone := -1
	dfs = func(r *TreeNode, f1 *int, f2 *int) int {
		if r == nil {
			return -1
		}
		if r.Val == *f1 {
			findone = *f1
			*f1 = -1
			return -1
		}
		if r.Val == *f2 {
			findone = *f2
			*f2 = -1
			return -1
		}
		L := dfs(r.Left, f1, f2)
		R := dfs(r.Right, f1, f2)
		if L != -1 {
			return L
		}
		if R != -1 {
			return R
		}
		if *f1 == -1 && *f2 == -1 {
			return r.Val
		}
		return -1
	}
	res := dfs(root, &o1, &o2)
	if res == -1 {
		return findone
	}
	return res
}

func pruneLeaves(root *TreeNode) *TreeNode {
	// write code here
	if root == nil {
		return nil
	}

	if isLeave(root.Left) || isLeave(root.Right) {
		return nil
	}

	root.Left = pruneLeaves(root.Left)
	root.Right = pruneLeaves(root.Right)
	return root
}

func numKLenSubstrRepeats(s string, k int) int {
	// write code here
	rem := make([]int, 26)
	ans := 0
	for i := 0; i < len(s); i++ {
		rem[s[i]-'a']++
		if i >= k {
			for j := 0; j < 26; j++ {
				if rem[j] >= 2 {
					ans++
					break
				}
			}
			rem[s[i-k]-'a']--
		}
	}
	return ans
}

func UniquePerm(nums []int) [][]int {
	// write code here
	sort.Ints(nums)
	chosed := make([]bool, len(nums))
	snums := make([]int, len(nums))

	var res [][]int
	var dfs = func(i int) {}
	dfs = func(i int) {
		if i == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}
		before := -1
		for j := 0; j < len(nums); j++ {
			if !chosed[j] {
				if before == -1 || nums[j] != nums[before] {
					before = j
					snums[i] = nums[j]
					chosed[j] = true
					dfs(i + 1)
					chosed[j] = false
				}
			}
		}
	}
	dfs(0)
	return res
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// write code here
	head := &ListNode{}
	p := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return head.Next
}
func isLeave(p *TreeNode) bool {
	return p != nil && p.Left == nil && p.Right == nil
}

func maxIntValue(arrs []int) int {
	// write code here
	sort.Slice(arrs, func(i, j int) bool {
		return arrs[j] > arrs[i]
	})
	res := 0
	for i := 0; i < len(arrs); i++ {
		res *= 10
		res += arrs[i]
	}
	return res
}

func replaceStr(s string) string {
	// write code here
	var res []byte
	has := false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			has = true
		} else {
			if has {
				res = append(res, []byte("num")...)
				has = false
			}
			res = append(res, s[i])
		}
	}
	if has {
		res = append(res, []byte("num")...)
	}
	return string(res)
}

func numIslands(grid [][]string) int {
	n := len(grid)
	m := len(grid[0])
	// write code here
	steps := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	res := 0
	dfs := func(i, j int) {}
	dfs = func(i, j int) {
		grid[i][j] = "0"
		for _, step := range steps {
			nextI := i + step[0]
			nextJ := j + step[1]
			if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m && grid[nextI][nextJ] == "1" {
				dfs(nextI, nextJ)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == "1" {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

func main() {
	a, b, c, d, e := 0, 0, 0, 0, 0
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	fmt.Scanln(&e)
	ans := 0
	if a < 0 && b > 0 {
		ans = (-a * c) + d + (b * e)
	} else if a > 0 {
		ans = (b - a) * e
	} else {
		ans = (b - a) * c
	}

	fmt.Print(ans)
}
