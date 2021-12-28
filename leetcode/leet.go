package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math"
	"math/bits"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func max(args ...int) int {
	ans := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > ans {
			ans = args[i]
		}
	}
	return ans
}

func min(args ...int) int {
	ans := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < ans {
			ans = args[i]
		}
	}
	return ans
}

// Definition for Employee.
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func isRectangleCover(rectangles [][]int) bool {
	area := 0
	MAX := 100005
	MIN := -100005
	a1, b1, a2, b2 := MAX, MAX, MIN, MIN
	set := make(map[string]int)
	for _, rectangle := range rectangles {
		x1, y1, x2, y2 := rectangle[0], rectangle[1], rectangle[2], rectangle[3]
		area += (x2 - x1) * (y2 - y1)
		if x1 < a1 || y1 < b1 {
			a1 = x1
			b1 = y1
		}
		if x2 > a2 || y2 > b2 {
			a2 = x2
			b2 = y2
		}
		// 记录每个顶点出现的次数
		record(set, x1, y1)
		record(set, x1, y2)
		record(set, x2, y1)
		record(set, x2, y2)
	}
	total := (a2 - a1) * (b2 - b1)
	if total != area {
		return false
	}
	return len(set) == 4 && has(set, a1, b1) && has(set, a1, b2) && has(set, a2, b1) && has(set, a2, b2)

}

func has(set map[string]int, a int, b int) bool {
	key := strconv.Itoa(a) + "-" + strconv.Itoa(b)
	_, ok := set[key]
	return ok
}

func record(set map[string]int, x1 int, y1 int) {
	key := strconv.Itoa(x1) + "-" + strconv.Itoa(y1)
	set[key]++
	if set[key]%2 == 0 {
		delete(set, key)
	}
}

func kInversePairs(n int, k int) int {
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}
	mod := int(1e9 + 7)
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		dp[i][0] = 1
		for j := 1; j <= k; j++ {
			if j-i-1 >= 0 {
				dp[i][j] -= dp[i-1][j-i-1]
			}
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
			dp[i][j] %= mod
		}
	}
	return dp[n-1][k]
}

func findMinStep(board string, hand string) int {
	var rem [26]int
	Co := "RYBGW"
	for _, v := range hand {
		rem[v-'A']++
	}
	dp := make(map[string]int)

	var dfs func(string, int) int
	dfs = func(board string, step int) int {
		if board == "" {
			return step
		}
		if v, has := dp[board]; has {
			return v
		}
		ans := 6
		for k, C := range Co {
			if rem[C-'A'] > 0 {
				rem[C-'A']--
				for i := 0; i <= len(board); i++ {
					ans = min(dfs(inrow(board[0:i]+Co[k:k+1]+board[i:]), step+1), ans)
				}
				rem[C-'A']++
			}
		}
		dp[board] = ans
		return ans
	}
	ans := dfs(board, 0)
	if ans == 6 {
		return -1
	}
	return ans
}

func inrow(s string) string {
	if len(s) < 3 {
		return s
	}
	l := 0
	for l < len(s) {
		r := l + 1
		for r < len(s) && s[l] == s[r] {
			r++
		}
		if r-l >= 3 {
			return inrow(s[0:l] + s[r:])
		}
		l = r
	}
	return s
}

func findWords(words []string) []string {
	ss := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	m := make([]int, 26)
	for i := 0; i < len(ss); i++ {
		for _, s := range ss[i] {
			m[s-'a'] = i
		}
	}
	var res []string
	for _, word := range words {
		i := 0
		word = strings.ToLower(word)
		for _, s := range word {

			if m[s-'a'] != m[word[0]-'a'] {
				break
			}
			i++
		}
		if i == len(word) {
			res = append(res, word)
		}
	}
	return res
}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	rem := make(map[int]int)
	var stack []int
	for _, num := range nums2 {
		for len(stack) != 0 && stack[len(stack)-1] < num {
			rem[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	var res []int
	for i := 0; i < len(nums1); i++ {
		if v, ok := rem[nums1[i]]; ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}

	}
	return res
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	dp := make(map[string]int)
	var dfs func(need []byte) int
	dfs = func(need []byte) (minPrice int) {
		if res, ok := dp[string(need)]; ok {
			return res
		}
		for i, p := range price {
			minPrice += int(need[i]) * p // 不购买任何大礼包，原价购买购物清单中的所有物品
		}
		next := make([]byte, n)
	out:
		for _, s := range special {
			for i, nd := range need {
				if int(nd) < s[i] {
					continue out
				}
				next[i] = need[i] - byte(s[i])
			}
			minPrice = min(minPrice, dfs(next))
		}
		return
	}
	need := make([]byte, n)
	for i, nd := range needs {
		need[i] = byte(nd)
	}
	return dfs(need)
}
func fib(n int) int {
	const M int = 1e9 + 7
	if n < 2 {
		return n
	}
	fn1, fn2 := 0, 1
	for i := 1; i < n; i++ {
		fn1, fn2 = fn2, (fn2+fn1)%M
	}
	return fn2
}

//func addBinary(a string, b string) string {
//	strconv.ParseInt(a,b)
//}
func reorderedPowerOf2(n int) bool {
	tmp := strconv.Itoa(n)
	nums := []byte(tmp)

	rem := make(map[string]bool)
	prefix := make(map[string]bool)
	for i := 0; i < 32; i++ {
		tmp := strconv.Itoa(1 << i)
		rem[tmp] = true
		for i2, _ := range tmp {
			prefix[tmp[:i2+1]] = true
		}
	}
	var res bool
	var dfs func([]byte, int, *bool)
	dfs = func(nums []byte, i int, res *bool) {
		if i == len(nums)-1 {
			if rem[string(nums)] {
				*res = true
			}
			return
		}
		if !*res {
			for j := i; j < len(nums); j++ {
				if i == 0 && nums[j] == 0 {
					continue
				}

				nums[i], nums[j] = nums[j], nums[i]
				if prefix[string(nums[:i+1])] {
					dfs(nums, i+1, res)
				}
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	dfs(nums, 0, &res)
	return res
}

func removeInvalidParentheses(s string) []string {
	lremove, rremove := 0, 0
	for _, i := range s {
		if i == '(' {
			lremove++
		} else if i == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}
	res := make(map[string]struct{})
	helper(s, 0, "", lremove, rremove, 0, 0, &res)
	var r []string
	for key, _ := range res {
		r = append(r, key)
	}
	return r
}

func helper(s string, i int, tmp string, lremove int, rremove int, lcount int, rcout int, res *map[string]struct{}) {
	if lremove == 0 && rremove == 0 {
		(*res)[tmp+s[i:]] = struct{}{}
		return
	}
	if s[i] != '(' && s[i] != ')' {
		helper(s, i+1, tmp+s[i:i+1], lremove, rremove, lcount, rcout, res)
	} else {
		if s[i] == '(' && lremove > 0 {
			helper(s, i+1, tmp, lremove-1, rremove, lcount, rcout, res)
		}
		if s[i] == ')' && rremove > 0 {
			helper(s, i+1, tmp, lremove, rremove-1, lcount, rcout, res)
		}
		if lremove+rremove <= len(s)-i-1 {
			if s[i] == ')' && rcout+1 > lcount {
				return
			}
			if s[i] == ')' {
				helper(s, i+1, tmp+s[i:i+1], lremove, rremove, lcount, rcout+1, res)
			} else {
				helper(s, i+1, tmp+s[i:i+1], lremove, rremove, lcount+1, rcout, res)
			}
		}
	}
}

func numDecodings(s string) int {
	if s == "0" {
		return 0
	}
	MAX := int(1e9 + 7)
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] == '*' {
		dp[1] = 9
	} else {
		dp[1] = 1
	}
	for i := 1; i < len(s); i++ {
		be := s[i-1]
		if s[i] == '*' {
			dp[i+1] = (9 * dp[i]) % MAX
			if be == '1' {
				dp[i+1] = (dp[i+1] + dp[i-1]*9) % MAX
			}
			if be == '2' {
				dp[i+1] = (dp[i+1] + dp[i-1]*6) % MAX
			}
			if be == '*' {
				dp[i+1] = (dp[i+1] + dp[i-1]*15) % MAX
			}
		} else if s[i] == '0' {
			if be == '*' {
				dp[i+1] = 2 * dp[i-1] % MAX
			} else if be == '1' || be == '2' {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else {
			dp[i+1] = dp[i]
			if be <= '6' {
				if be == '*' {
					dp[i+1] = (dp[i+1] + 2*dp[i-1]) % MAX
				}
				if be == '1' || be == '2' {
					dp[i+1] = (dp[i+1] + dp[i-1]) % MAX
				}
			} else {
				if be == '*' || be == '1' {
					dp[i+1] = (dp[i+1] + dp[i-1]) % MAX
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)]
}

func divide(a int, b int) int {
	var flag int
	if a^b > 0 {
		flag = 1
	} else {
		flag = -1
	}
	a = abs(a)
	b = abs(b)
	res := 0

	for a > 0 {
		shirt := 0
		for (b << shirt) <= a {
			shirt++
		}
		a -= b << (shirt - 1)
		res += 1 << (shirt - 1)
	}
	return res * flag
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	wl := len(words[0])
	start := 0
	for i := 1; i < len(words); i++ {
		if wl+len(words[i])+i-start <= maxWidth {
			wl += len(words[i])
		} else {
			res = append(res, buildLine(wl, words[start:i], maxWidth))
			wl = len(words[i])
			start = i
		}
	}
	var line string
	for i := start; i < len(words); i++ {
		line += words[i]
		if i != len(words)-1 {
			line += " "
		}
	}
	line += strings.Repeat(" ", maxWidth-len(line))
	res = append(res, line)
	return res
}

func buildLine(wl int, words []string, maxWidth int) string {
	wn := len(words)
	if wn == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-len(words[0]))
	}
	var line string
	p := (maxWidth - wl) / (wn - 1)
	less := maxWidth - wl - p*(wn-1)
	gap := strings.Repeat(" ", p)
	for i := 0; i < wn; i++ {
		line += words[i]
		if i != wn-1 {
			line += gap
			if less > 0 {
				line += " "
				less--
			}
		}
	}
	return line
}

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	l, r := 0, len(arr)-1
	for l < r {
		p := patition(arr, l, r)
		if p == k {
			return arr[:k]
		} else if p > k {
			r = p - 1
		} else {
			k = p + 1
		}
	}
	return arr
}

func patition(arr []int, l, r int) int {
	p := arr[l]
	for l < r {
		for r > l && arr[r] >= p {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= p {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = p
	return p
}

func patition2(arr []int, l, r int) (int, int) {
	p := arr[l]
	i := l + 1
	for i <= r {
		if arr[i] > p {
			arr[i], arr[r] = arr[r], arr[i]
			r--
		} else if arr[i] < p {
			arr[i], arr[l] = arr[l], arr[i]
			i++
			l++
		} else {
			i++
		}
	}
	return l, r
}

func compress(chars []byte) int {
	num := 0
	t := 0
	for i := 0; i < len(chars); i++ {
		num++
		if i == len(chars)-1 || chars[i] != chars[i+1] {
			chars[t] = chars[i]
			t++
			if num > 1 {
				tmp := strconv.Itoa(num)
				for n := 0; n < len(tmp); n++ {
					chars[t] = tmp[n]
					t++
				}
			}
			num = 0
		}
	}
	return t
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
//	return ls
//}

func eventualSafeNodes(graph [][]int) (res []int) {

	color := make([]int, len(graph))
	var safe func(int) bool
	safe = func(x int) bool {
		if color[x] > 0 {
			return color[x] == 2
		}
		color[x] = 1
		for i := 0; i < len(graph[x]); i++ {
			if !safe(graph[x][i]) {
				return false
			}
		}
		color[x] = 2
		return true
	}
	for i := 0; i < len(graph); i++ {
		if safe(i) {
			res = append(res, i)
		}
	}
	return
}

func findUnsortedSubarray(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(nums)
	i := 0
	for ; i < len(nums); i++ {
		if tmp[i] != nums[i] {
			break
		}
	}
	j := len(nums) - 1
	for ; j > i; j-- {
		if tmp[j] != nums[j] {
			break
		}
	}
	return j - i + 1
}
func findUnsortedSubarray2(nums []int) int {
	l, r := -1, -1
	maxn := math.MinInt32
	minn := math.MaxInt32
	n := len(nums)
	for i := 0; i < n; i++ {
		if maxn < nums[i] {
			maxn = nums[i]
		} else {
			r = i
		}
		if minn > nums[n-i-1] {
			minn = nums[n-i-1]
		} else {
			l = n - i - 1
		}
	}
	if r == -1 {
		return 0
	}
	return r - l + 1
}
func networkDelayTime(times [][]int, n int, k int) int {
	G := make([][]int, n)
	for i := 0; i < n; i++ {
		G[i] = make([]int, n)
		for h := 0; h < n; h++ {
			G[i][h] = math.MaxInt32
		}
	}
	for i := 0; i < len(times); i++ {
		G[times[i][0]-1][times[i][1]-1] = times[i][2]
	}
	dist := make([]int, n)
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k-1] = 0
	for i := 0; i < n-1; i++ {
		m := math.MaxInt32
		t := 0
		for j := 0; j < n; j++ {
			if dist[j] < m && !vis[j] {
				m = dist[j]
				t = j
			}
		}
		if m == math.MaxInt32 {
			break
		}
		vis[t] = true
		for j := 0; j < n; j++ {
			if G[t][j] != 0 && G[t][j]+m < dist[j] {
				dist[j] = G[t][j] + m
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if dist[i] > res {
			res = dist[i]
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

type node struct {
	row, col, val int
}

var nodes []node

func verticalTraversal(root *TreeNode) (res [][]int) {
	nodes = []node{}
	nodes = append(nodes, node{0, 0, root.Val})
	dfs987(root, 0, 0)
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].col == nodes[i].col && nodes[i].row == nodes[j].row {
			return nodes[i].val < nodes[j].val
		}
		if nodes[i].col == nodes[j].col {
			return nodes[i].row < nodes[j].row
		}
		return nodes[i].col < nodes[j].col
	})
	res = append(res, []int{nodes[0].val})
	for i := 1; i < len(nodes); i++ {
		if nodes[i].col != nodes[i-1].col {
			res = append(res, []int{nodes[i].val})
		} else {
			res[len(res)-1] = append(res[len(res)-1], nodes[i].val)
		}
	}

	return
}
func dfs987(root *TreeNode, row, col int) {
	if root.Left != nil {
		nodes = append(nodes, node{row + 1, col - 1, root.Left.Val})
		dfs987(root.Left, row+1, col-1)
	}

	if root.Right != nil {
		nodes = append(nodes, node{row + 1, col + 1, root.Right.Val})
		dfs987(root.Right, row+1, col+1)
	}
}

var res863 []int

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	res863 = []int{}

	do863(root, target, k)
	return res863
}

func do863(root *TreeNode, target *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	if root == target {
		dfs863(root, 0, k)
		return 1
	} else {
		r := do863(root.Right, target, k)
		if r > 0 {
			if r == k {
				res863 = append(res863, root.Val)
				return 0
			}
			if r < k {
				dfs863(root.Left, r+1, k)
			}
			return r + 1
		} else {
			l := do863(root.Left, target, k)
			if l > 0 {
				if l == k {
					res863 = append(res863, root.Val)
					return 0
				}
				if l < k {
					dfs863(root.Right, l+1, k)
				}
				return l + 1
			}
		}
	}
	return 0
}

func dfs863(root *TreeNode, tmp int, tar int) {
	if root == nil {
		return
	}
	if tmp == tar {
		res863 = append(res863, root.Val)
		return
	} else {
		dfs863(root.Left, tmp+1, tar)
		dfs863(root.Right, tmp+1, tar)
	}
}
func maximumTime(time string) string {
	tmp := []byte(time)
	if tmp[0] == '?' {
		if tmp[1] == '?' || tmp[1] < '4' {
			tmp[0] = '2'
		} else {
			tmp[0] = '1'
		}
	}
	if tmp[1] == '?' {
		if tmp[0] == '2' {
			tmp[1] = '3'
		} else {
			tmp[1] = '9'
		}
	}
	if tmp[3] == '?' {
		tmp[3] = '5'
	}
	if tmp[4] == '?' {
		tmp[4] = '9'
	}
	return string(tmp)
}

func groupAnagrams(strs []string) [][]string {
	rem := make(map[string]int)
	var ans [][]string
	for i := 0; i < len(strs); i++ {
		s := stringsort(strs[i])
		if p, ok := rem[s]; ok {
			ans[p] = append(ans[p], strs[i])
		} else {
			tmp := []string{strs[i]}
			ans = append(ans, tmp)
			rem[s] = len(ans) - 1
		}
	}
	return ans
}

func stringsort(s string) string {
	var nums [26]int
	for i := 0; i < len(s); i++ {
		nums[s[i]-'a']++
	}
	var res []byte
	for i := 0; i < 26; i++ {
		for j := 0; j < nums[i]; j++ {
			res = append(res, byte(i+'a'))
		}
	}
	return string(res)

	//t := []byte(s)
	//sort.Slice(t, func(i, j int) bool {
	//	return t[i] < t[j]
	//})
	//return string(t)
}
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	rem := make([]int, len(arr)+1)
	num := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] <= len(arr) {
			rem[arr[i]]++
		} else {
			num++
		}
	}
	ans := 0
	for i := 1; i <= len(arr); i++ {
		ans = min(i, ans+rem[i])
	}
	ans += num
	return ans
}

func numWays(n int, relation [][]int, k int) int {
	path := make([][]int, 10)
	for _, rela := range relation {
		path[rela[0]] = append(path[rela[0]], rela[1])
	}
	var ans int
	dfsnumWays(path, 0, 0, k, n-1, &ans)
	return ans
}

func dfsnumWays(path [][]int, tk int, tn int, k int, n int, ans *int) {
	if tk == k {
		if tn == n {
			*ans++
		}
		return
	}
	for i := 0; i < len(path[tn]); i++ {
		dfsnumWays(path, tk+1, path[tn][i], k, n, ans)
	}
}
func convertToTitle(columnNumber int) string {
	var ans []byte
	for columnNumber > 0 {
		tmp := (columnNumber - 1) % 26
		ans = append([]byte{byte(tmp + 'A')}, ans...)
		columnNumber -= tmp + 1
		columnNumber /= 26
	}
	return string(ans)
}
func snakesAndLadders(board [][]int) int {
	lr := len(board)
	lc := len(board[0])
	end := lr * lc
	points := make([]int, end+1)
	flag := 0
	p := 1
	for i := lr - 1; i >= 0; i-- {
		if flag == 0 {
			flag = 1
			for j := 0; j < lc; j++ {
				points[p] = board[i][j]
				p++
			}
		} else {
			flag = 0
			for j := lc - 1; j >= 0; j-- {
				points[p] = board[i][j]
				p++
			}
		}

	}
	step := make([]int, end+1)
	step[1] = 1
	var queue []int
	queue = append(queue, 1)
	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		if tmp == end {
			return step[tmp] - 1
		}
		for next := tmp + 1; next < tmp+7 && next <= end; next++ {
			if points[next] != -1 {
				if step[points[next]] == 0 {
					step[points[next]] = step[tmp] + 1
					queue = append(queue, points[next])
				}
			} else {
				if step[next] == 0 {
					step[next] = step[tmp] + 1
					queue = append(queue, next)
				}
			}
		}
	}
	return -1
}
func slidingPuzzle(board [][]int) int {
	var queue []string
	begin := ""
	end := "123450"
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			begin += strconv.Itoa(board[i][j])
		}
	}
	visit := make(map[string]int)
	visit[begin] = 0
	queue = append(queue, begin)
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p == end {
			return visit[p]
		}
		nexts := getNexts([]byte(p))
		for i := 0; i < len(nexts); i++ {
			if _, ok := visit[nexts[i]]; !ok {
				visit[nexts[i]] = visit[p] + 1
				queue = append(queue, nexts[i])
			}
		}
	}
	return -1
}

func getNexts(p []byte) []string {
	var place int
	for i := 0; i < 6; i++ {
		if p[i] == '0' {
			place = i
			break
		}
	}
	switch place {
	case 0:
		return []string{swaps(p, 0, 1), swaps(p, 0, 3)}
	case 1:
		return []string{swaps(p, 1, 0), swaps(p, 1, 2), swaps(p, 1, 4)}
	case 2:
		return []string{swaps(p, 2, 1), swaps(p, 2, 5)}
	case 3:
		return []string{swaps(p, 3, 0), swaps(p, 3, 4)}
	case 4:
		return []string{swaps(p, 4, 3), swaps(p, 4, 5), swaps(p, 4, 1)}
	default:
		return []string{swaps(p, 5, 4), swaps(p, 5, 2)}
	}
}
func swaps(tmp []byte, i, j int) string {
	tmp[i], tmp[j] = tmp[j], tmp[i]
	res := string(tmp)
	tmp[i], tmp[j] = tmp[j], tmp[i]
	return res
}

func openLock(deadends []string, target string) int {
	var queue []string
	var visted [10000]int
	for _, deadend := range deadends {
		i, _ := strconv.Atoi(deadend)
		visted[i] = -1
	}
	if visted[0] == 0 {
		queue = append(queue, "0000")
	}
	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		before, _ := strconv.Atoi(tmp)
		if tmp == target {
			return visted[before]
		}
		for i := 0; i < len(tmp); i++ {
			var next1, next2 byte
			if tmp[i] == '9' {
				next1 = 0
			} else {
				next1 = tmp[i] + 1 - '0'
			}
			if tmp[i] == '0' {
				next2 = 9
			} else {
				next2 = tmp[i] - 1 - '0'
			}
			nexts1 := tmp[:i] + strconv.Itoa(int(next1)) + tmp[i+1:]
			nexts2 := tmp[:i] + strconv.Itoa(int(next2)) + tmp[i+1:]
			i1, _ := strconv.Atoi(nexts1)
			i2, _ := strconv.Atoi(nexts2)
			if visted[i1] == 0 {
				visted[i1] = visted[before] + 1
				queue = append(queue, nexts1)
			}
			if visted[i2] == 0 {
				visted[i2] = visted[before] + 1
				queue = append(queue, nexts2)
			}
		}
	}
	return -1
}

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

func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	for i := 0; i < len(ranges); i++ {
		if ranges[i][0] <= left && ranges[i][1] >= left {
			left = ranges[i][1]
		}
	}
	return left >= right
}

func preOder(name string, nodes map[string][]string, res *[]string) {
	if _, ok := deathed[name]; !ok {
		*res = append(*res, name)
	}
	for _, son := range nodes[name] {
		preOder(son, nodes, res)
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

func sampleStats(count []int) []float64 {
	sum := 0
	n := 0
	for i, nu := range count {
		n += nu
		sum += i * nu
	}
	min, max, avg, mid, mos := -1.0, -1.0, float64(sum)/float64(n), -1.0, -1.0
	one := -1.0
	two := -1.0
	num := 0
	num_m := 0
	for i := 0; i < 256; i++ {
		t := float64(i)
		if count[i] != 0 {
			num += count[i]
			if min == -1.0 {
				min = t
			}
			if float64(i) > max {
				max = t
			}
			if num_m < count[i] {
				mos = t
				num_m = count[i]
			}
		}
		if n&1 == 0 {
			if num >= n/2 && one == -1 {
				one = t
			}
			if num >= n/2+1 && two == -1 {
				two = t
			}
		}
		if n&1 == 1 {
			if num >= n/2+1 && one == -1 {
				one = t
			}
		}
	}
	if n&1 == 1 {
		mid = one
	} else {
		mid = (one + two) / 2
	}

	return []float64{min, max, avg, mid, mos}
}
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	dnum := 0
	a, b := -1, -1
	rem := make([]int, 26)
	two := false
	for i := 0; i < len(s); i++ {
		if !two {
			rem[s[i]-'a']++
			if rem[s[i]-'a'] == 2 {
				two = true
			}
		}
		if s[i] != goal[i] {
			dnum++
			if a == -1 {
				a = i
			} else if b == -1 {
				b = i
			}
		}
	}
	if dnum == 2 {
		return s[a] == goal[b] && s[b] == goal[a]
	}
	if dnum == 0 {
		return two
	}
	return false
}

func findAnagrams(s string, p string) []int {
	var mask [26]int
	n := len(p)
	for i := 0; i < n; i++ {
		mask[p[i]-'a']++
	}
	gap := n
	var ans []int
	for i := 0; i < len(s); i++ {
		mask[s[i]-'a']--
		if mask[s[i]-'a'] >= 0 {
			gap--
		} else {
			gap++
		}
		if gap == 0 {
			ans = append(ans, i-n+1)
		}
		if i >= n-1 {
			mask[s[i-n+1]-'a']++
			if mask[s[i-n+1]-'a'] > 0 {
				gap++
			} else {
				gap--
			}
		}
	}
	return ans
}

type node2 struct {
	a, b int
	c    float64
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	var ns []float32
	rem := make(map[float32][]int)
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr); j++ {
			f := float32(arr[i]) / float32(arr[j])
			ns = append(ns, f)
			rem[f] = []int{arr[i], arr[j]}
		}
	}
	kselect(k-1, ns, 0, len(ns)-1)
	return rem[ns[k-1]]
}

func kselect(k int, ns []float32, l, r int) {
	for l < r {
		p := partition(ns, l, r)
		if p == k {
			return
		} else if p > k {
			r = p - 1
		} else {
			l = p + 1
		}
	}
}

func partition(ns []float32, i, j int) int {
	x := ns[i]
	for i < j {
		for j > i && ns[j] >= x {
			j--
		}
		ns[i] = ns[j]
		for i < j && ns[i] <= x {
			i++
		}
		ns[j] = ns[i]
	}
	ns[i] = x
	return i
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	m := 101
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
		if nums[i] < m {
			m = nums[i]
		}
		sum += nums[i]
	}
	if k > 0 && k%2 == 1 {
		return sum - 2*m
	}
	return sum
}

type name struct {
	a       int
	bdsfadf string
}

type TopVotedCandidate struct {
	tops, times []int
}

func TConstructor(persons, times []int) TopVotedCandidate {
	tops := make([]int, len(persons))
	top := -1
	voteCounts := map[int]int{-1: -1}
	for i, p := range persons {
		voteCounts[p]++
		if voteCounts[p] >= voteCounts[top] {
			top = p
		}
		tops[i] = top
	}

	return TopVotedCandidate{tops, times}
}
func (c *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(c.times)
	t++
	for l < r {
		m := (l + r) >> 1
		if c.times[m] >= t {
			r = m
		} else {
			l = m + 1
		}
	}
	return c.tops[l-1]
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][0] == courses[j][0] {
			return courses[i][1] < courses[j][1]
		}
		return courses[i][0] < courses[j][0]
	})

	ans := 0
	last := 0
	for i := 0; i < len(courses); i++ {
		if courses[i][1]-courses[i][0] >= last {
			ans++
			last += courses[i][0]
		}
	}
	return ans
}

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	G := make([][]int, n)
	for _, ints := range richer {
		G[ints[1]] = append(G[ints[1]], ints[0])
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	var dfs func(k int)
	dfs = func(k int) {
		ans[k] = k
		for i := 0; i < len(G[k]); i++ {
			if ans[G[k][i]] == -1 {
				dfs(G[k][i])
			}
			if quiet[ans[k]] > quiet[G[k][i]] {
				ans[k] = G[k][i]
			}
		}
	}
	for i := 0; i < n; i++ {
		if ans[i] == -1 {
			dfs(i)
		}
	}
	return ans
}

func longestDupSubstring(s string) string {
	l, r := 1, len(s)-1
	var ans string
	var has bool
	for l < r {
		m := (l + r) / 2
		set := make(map[string]struct{})
		has = false
		for i := 0; i <= len(s)-m; i++ {
			if _, has = set[s[i:i+m]]; has {
				ans = s[i : i+m]
				break
			}
		}
		if has {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}

func isEvenOddTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	for level := 0; len(q) > 0; level++ {
		prev := 0
		if level%2 == 1 {
			prev = math.MaxInt32
		}
		size := len(q)
		for _, node := range q {
			val := node.Val
			if val%2 == level%2 || level%2 == 0 && val <= prev || level%2 == 1 && val >= prev {
				return false
			}
			prev = val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
	}
	return true
}

func getUser(c echo.Context) error {
	//var u user
	//id := c.Param("id")
	//if !bson.IsObjectIdHex(id) {
	//	return newHTTPError(http.StatusBadRequest, "InvalidID", "invalid user id")
	//}
	//err := db.C("user").FindId(bson.ObjectIdHex(id)).One(&u)
	//if err == mgo.ErrNotFound {
	//	return newHTTPError(http.StatusNotFound, "NotFound", err.Error())
	//}
	//if err != nil {
	//	return err
	//}
	//return c.JSON(http.StatusOK, u)
	return nil
}

func numFriendRequests(ages []int) int {
	rem := make([]int, 125)
	for i := 0; i < len(ages); i++ {
		rem[ages[i]]++
	}
	for i := 1; i <= 120; i++ {
		rem[i] += rem[i-1]
	}
	ans := 0
	for i := 15; i <= 120; i++ {
		if rem[i] != 0 {
			bound := i/2 + 7
			ans += (rem[i] - rem[i-1]) * (rem[i] - rem[bound] - 1)
		}
	}
	return ans
}

func main() {

	wg := sync.WaitGroup{}
	go func(wg *sync.WaitGroup) {
		wg.Done()
	}(&wg)
}
