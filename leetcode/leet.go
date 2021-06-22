package main

import (
	"math"
	"math/bits"
	"reflect"
	"sort"
	"strconv"
)

// Definition for Employee.
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

//func firstBadVersion(n int) int {
//	l, r := 1, n
//	for l < r {
//		m := (l + r) / 2
//		if isBadVersion(m) {
//			r = m
//		} else {
//			l = m + 1
//		}
//	}
//	return l
//}

type ThroneInheritance struct {
	kingName string
	nodes    map[string][]string
}

var deathed map[string]struct{}

//func Constructor(kingName string) ThroneInheritance {
//	deathed = make(map[string]struct{})
//	tmp := ThroneInheritance{kingName: kingName, nodes: make(map[string][]string)}
//	return tmp
//}
//
//func (this *ThroneInheritance) Birth(parentName string, childName string) {
//	this.nodes[parentName] = append(this.nodes[parentName], childName)
//}
//
//func (this *ThroneInheritance) Death(name string) {
//	deathed[name] = struct{}{}
//}
//
//func (this *ThroneInheritance) GetInheritanceOrder() []string {
//	res := []string{}
//	preOder(this.kingName, this.nodes, &res)
//	return res
//}

func preOder(name string, nodes map[string][]string, res *[]string) {
	if _, ok := deathed[name]; !ok {
		*res = append(*res, name)
	}
	for _, son := range nodes[name] {
		preOder(son,nodes,res)
	}
}



func maxLength(arr []string) int {
	ans := 0
	var nums []uint32
	var lens []int
	lens = append(lens, 0)
	for i := 0; i < len(arr); i++ {
		var t uint32
		app := true
		for j := 0; j < len(arr[i]); j++ {
			if (t & (1 << (arr[i][j] - 'a'))) != 0 {
				app = false
				break
			}
			t |= 1 << (arr[i][j] - 'a')
		}
		if app {
			nums = append(nums, t)
			lens = append(lens, lens[len(nums)]+len(arr[i]))
		}
	}

	dfs1239(nums, lens, 0, 0, &ans, 0)
	return ans
}

func dfs1239(arr []uint32, lens []int, k int, tmp int, ans *int, selected uint32) {
	if k == len(arr) {
		if tmp > *ans {
			*ans = tmp
		}
		return
	}
	if (selected & arr[k]) == 0 {
		dfs1239(arr, lens, k+1, tmp+lens[k+1]-lens[k], ans, selected|arr[k])
	}
	if tmp+lens[len(arr)]-lens[k+1] > *ans {
		dfs1239(arr, lens, k+1, tmp, ans, selected)
	}
}

func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)
	maxl := bits.Len(uint(num)) - 1
	for i := maxl; i > 1; i-- {
		k := int(math.Pow(float64(num), float64(i)))
		sum := 1
		mul := 1
		for j := 0; j < i; j++ {
			mul *= k
			sum += mul
		}
		if sum == num {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(num - 1)
}

// 2021  6.6
func findMaxForm(strs []string, m int, n int) int {

	num0 := make([]int, len(strs))
	num1 := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		n0, n1 := 0, 0
		for _, c := range strs[i] {
			if c == '1' {
				n1++
			} else {
				n0++
			}
		}
		num0[i] = n0
		num1[i] = n1
	}
	dp := [101][101]int{}
	for i := 0; i < len(num0); i++ {
		for j := m; j >= num0[i]; j-- {
			for k := n; k >= num1[i]; k-- {
				dp[j][k] = max(dp[j][k], dp[j-num0[i]][k-num1[i]]+1)
			}
		}
	}
	return dp[m][n]
}

// 6.3   找到零和一数量相同的最长连续子数组长度
func findMaxLength(nums []int) int {
	num1, num0 := 0, 0
	rem := make(map[int]int)
	rem[0] = -1
	ans := 0
	for i := 0; i < len(nums); i++ {
		num1 += nums[i]
		num0 += 1 - nums[i]
		nums[i] = num1 - num0
		if p, ok := rem[nums[i]]; !ok {
			rem[nums[i]] = i
		} else {
			ans = max(i-p, ans)
		}
	}
	return ans
}

// Leetcode 690  5.1
func getImportance(employees []*Employee, id int) int {
	m := make(map[int]*Employee)
	for _, employee := range employees {
		m[employee.Id] = employee
	}
	return getImportanceByMap(m, id)
}

func getImportanceByMap(m map[int]*Employee, id int) int {
	tmp := m[id]
	var ans = tmp.Importance
	for _, subordinate := range tmp.Subordinates {
		ans += getImportanceByMap(m, subordinate)
	}
	return ans
}

// 554砖墙  2021 - 5.2
func leastBricks(wall [][]int) int {
	rem := make(map[int]int)
	for i := 0; i < len(wall); i++ {
		tmp := 0
		for j := 0; j < len(wall[i])-1; j++ {
			tmp += wall[i][j]
			rem[tmp]++
		}
	}
	var ans = len(wall)
	var maxSize = 0
	for _, size := range rem {
		if size > maxSize {
			maxSize = size
		}
	}
	return ans - maxSize
}

// 7 整数反转  2021 5.3
func reverse(x int) int {
	rev := 0
	for x != 0 {
		rev *= 10
		rev += x % 10
		x /= 10
	}
	if rev < math.MinInt32 || rev > math.MaxInt32 {
		return 0
	}
	return rev
}

// 1473粉刷房子 2021.5.4
func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	var dp [105][25][105]int
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j][0] = math.MaxInt32
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= target; k++ {
				if k > i {
					dp[i][j][k] = math.MaxInt32
					continue
				}
				if houses[i-1] != 0 && j != houses[i-1] {
					dp[i][j][k] = math.MaxInt32
				} else {
					dp[i][j][k] = dp[i-1][j][k]
					for p := 1; p <= n; p++ {
						if p != j {
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][p][k-1])
						}
					}
					if houses[i-1] == 0 {
						dp[i][j][k] += cost[i-1][j-1]
					}
				}
			}
		}
	}
	ans := math.MaxInt32
	for i := 1; i <= n; i++ {
		ans = min(ans, dp[m][i][target])
	}
	if ans != math.MaxInt32 {
		return ans
	}
	return -1
}

// 740删除并或得点数 2021.5.6
func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	dp := make([]int, len(nums)+1)
	var size []int
	tmp := 1
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 || nums[i] != nums[i+1] {
			size = append(size, tmp)
			nums[len(size)-1] = nums[i]
			tmp = 1
		} else {
			tmp++
		}
	}
	dp[1] = size[0] * nums[0]
	for i := 1; i < len(size); i++ {
		if nums[i] != nums[i-1]+1 {
			dp[i+1] = dp[i] + nums[i]*size[i]
		} else {
			dp[i+1] = max(dp[i], dp[i-1]+nums[i]*size[i])
		}
	}
	return dp[len(size)]
}

// 1720. 解码异或后的数组 2021.5.7
func Decode(encoded []int, first int) []int {
	var tmp int
	for i := 0; i < len(encoded); i++ {
		tmp = first
		first = encoded[i] ^ first
		encoded[i] = tmp
	}
	encoded = append(encoded, first)
	return encoded
}

// 1486. 数组异或操作  2021.5.7

// 1723. 完成所有工作的最短时间 2021.5.8
var ans1723 int

func minimumTimeRequired(jobs []int, k int) int {
	workTime := make([]int, k)
	ans1723 = math.MaxInt32
	dfs1723(0, jobs, 0, workTime)
	return ans1723
}

func dfs1723(i int, jobs []int, tmp int, workTime []int) {
	if i == len(jobs) {
		if tmp < ans1723 {
			ans1723 = tmp
		}
		return
	}
	sort.Ints(workTime)
	for j := 0; j < len(workTime); j++ {
		if workTime[j]+jobs[i] < ans1723 {
			if j > 0 && workTime[j] == workTime[j-1] {
				continue
			}
			workTime[j] += jobs[i]
			cop := make([]int, len(workTime))
			copy(cop, workTime)
			dfs1723(i+1, jobs, max(tmp, workTime[j]), cop)
			workTime[j] -= jobs[i]
		}
	}
}

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}
	r := 0
	l := math.MaxInt32
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] > r {
			r = bloomDay[i]
		}
		if bloomDay[i] < l {
			l = bloomDay[i]
		}
	}

	for l < r {
		m := (l + r) / 2
		tmp := 0
		sum := 0
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] <= m {
				tmp++
				if tmp == k {
					sum++
					tmp = 0
				}
			} else {
				tmp = 0
			}
		}
		if sum >= m {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := getLeaf(root1)
	leafs2 := getLeaf(root2)
	if len(leafs1) != len(leafs2) {
		return false
	}
	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}
	return true
}

func getLeaf(root *TreeNode) []int {
	var leafs []int
	if root != nil {
		dfs872(root, &leafs)
	}
	return leafs
}

func dfs872(root *TreeNode, i *[]int) {
	if root.Left == nil && root.Right == nil {
		*i = append(*i, root.Val)
		return
	}
	if root.Left != nil {
		dfs872(root.Left, i)
	}
	if root.Right != nil {
		dfs872(root.Right, i)
	}
}

type Dog struct {
	Id   int
	Name string
}

var m map[string]interface{}

func get(key string, tmp interface{}) {
	value := m[key]
	reflect.ValueOf(tmp).Elem().Set(reflect.ValueOf(value))
}

func decode(encoded []int) []int {
	oxN := 0
	for i := 1; i <= len(encoded)+1; i++ {
		oxN ^= i
	}
	first := 0
	for i := 1; i < len(encoded); i += 2 {
		first ^= encoded[i]
	}
	var ans []int
	ans = append(ans, first)
	for i := 0; i < len(encoded); i++ {
		ans = append(ans, ans[i]^encoded[i])
	}
	return ans
}

func findMaximumXOR1(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	x := 0
	for k := 30; k >= 0; k-- {
		m := make(map[int]struct{})
		for i := 0; i < len(nums); i++ {
			m[nums[i]>>k] = struct{}{}
		}
		x_next := 2*x + 1
		find := false
		for _, num := range nums {
			if _, ok := m[x_next^(num>>k)]; ok {
				find = true
				break
			}
		}
		if find {
			x = x_next
		} else {
			x = x_next - 1
		}
	}
	return x
}
func isCousins(root *TreeNode, x int, y int) bool {
	if root.Val == x || root.Val == y {
		return false
	}
	last := root
	var queue []*TreeNode
	queue = append(queue, root)
	var lx, ly, l int
	l = 1
	var px, py *TreeNode
	for len(queue) != 0 {
		root = queue[0]
		queue = queue[1:]
		if root.Left != nil {
			if root.Left.Val == x {
				lx = l
				px = root
			}
			if root.Left.Val == y {
				ly = l
				py = root
			}
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			if root.Right.Val == x {
				lx = l
				px = root
			}
			if root.Right.Val == y {
				ly = l
				py = root
			}
			queue = append(queue, root.Right)
		}
		if last == root {
			l++
			last = queue[len(queue)-1]
		}
		if lx != 0 && ly != 0 {
			break
		}
	}
	return lx == ly && px != py
}

func kthLargestValue(matrix [][]int, k int) int {
	nums := []int{}
	for i := 1; i < len(matrix); i++ {
		matrix[i][0] ^= matrix[i-1][0]
		nums = append(nums, matrix[i][0])
	}
	for i := 1; i < len(matrix[0]); i++ {
		matrix[0][i] ^= matrix[0][i-1]
		nums = append(nums, matrix[0][i])
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			matrix[i][j] ^= matrix[i-1][j-1] ^ matrix[i-1][j] ^ matrix[i][j-1]
			nums = append(nums, matrix[i][j])
		}
	}
	return quickSelect(nums, k)
}

func quickSelect(nums []int, k int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		i := l
		j := r
		val := nums[l]
		for i < j {
			for j > i && nums[j] <= val {
				j--
			}
			nums[i] = nums[j]
			i++
			for i < j && nums[i] >= val {
				i++
			}
			nums[j] = nums[i]
			j--
		}
		nums[i] = val
		if i+1 == k {
			break
		} else if i+1 > k {
			r = i - 1
		} else {
			l = i + 1
		}
	}
	return nums[k-1]
}

type stringTime struct {
	i   int
	str string
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	var strs []stringTime
	for s, i := range m {
		strs = append(strs, stringTime{i, s})
	}
	sort.Slice(strs, func(i, j int) bool {
		if strs[i].i == strs[j].i {
			return strs[i].str < strs[j].str
		}
		return strs[i].i > strs[j].i
	})
	var ans []string
	for i := 0; i < k; i++ {
		ans = append(ans, strs[i].str)
	}
	return ans
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i-1][j-1] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}

func main() {
}
