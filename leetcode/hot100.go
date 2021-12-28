package main

import "sort"

// 78. 子集
func subsets(nums []int) [][]int {
	var ans [][]int
	var dfs func(k int, tmp []int)
	dfs = func(k int, tmp []int) {
		if k == len(nums) {
			one := make([]int, len(tmp))
			copy(one, tmp)
			ans = append(ans, one)
			return
		}
		dfs(k+1, tmp)
		dfs(k+1, append(tmp, nums[k]))
	}
	dfs(0, []int{})
	return ans
}

// 79 单词搜索
func exist(board [][]byte, word string) bool {
	steps := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	n := len(board)
	m := len(board[0])
	var visit []int16
	find := false
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		if !find {
			if k == len(word)-1 {
				find = true
				return
			}
			for _, step := range steps {
				nexti := i + step[0]
				nextj := j + step[1]
				if nexti >= 0 && nexti < n && nextj >= 0 && nextj < m {
					if (visit[nexti]&(1<<nextj)) == 0 && board[nexti][nextj] == word[k+1] {
						visit[nexti] |= 1 << nextj
						dfs(nexti, nextj, k+1)
						visit[nexti] &= ^(1 << nextj)
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				visit = make([]int16, n)
				visit[i] |= 1 << j
				dfs(i, j, 0)
			}
		}
	}
	return find
}

func maxProduct(nums []int) int {
	Max, Min := nums[0], nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		Max = Max * nums[i]
		Min = Min * nums[i]
		Max, Min = max(Max, Min, nums[i]), min(Max, Min, nums[i]) // 不能分开写
		if Max > ans {
			ans = Max
		}
	}
	return ans
}

func rob(nums []int) int {
	ans := nums[0]
	b1 := nums[0]
	b2 := 0
	for i := 1; i < len(nums); i++ {
		b1, b2 = max(b2+nums[i], b1), b1
		ans = max(ans, b1)
	}
	return ans
}

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	ans := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1], dp[i][j]) + 1
				if dp[i+1][j+1] > ans {
					ans = dp[i+1][j+1]
				}
			}
		}
	}
	return ans * ans
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	ans := &TreeNode{}
	var dfs func(root, p, q *TreeNode) bool
	dfs = func(root, p, q *TreeNode) bool {
		if root != nil {
			l := dfs(root.Left, p, q)
			r := dfs(root.Right, p, q)
			if (l && r) || (root.Val == p.Val || root.Val == q.Val) && (l || r) {
				ans = root
			}
			return r || l || root.Val == q.Val || root.Val == p.Val
		}
		return false
	}
	dfs(root, p, q)
	return ans
}

func findAllConcatenatedWordsInADict(words []string) []string {
	if len(words) <= 2 {
		return []string{}
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	rem := make(map[string]struct{})
	rem[words[0]] = struct{}{}
	rem[words[1]] = struct{}{}
	var ans []string
	for i := 2; i < len(words); i++ {
		var cut_p = []int{0}
		for j := 0; j < len(words[i]); j++ {
			for c := 0; c < len(cut_p); c++ {
				if _, has := rem[words[i][cut_p[c]:j+1]]; has {
					cut_p = append(cut_p, j+1)
					break
				}
			}
			if cut_p[len(cut_p)-1] == j+1 {
				if _, has := rem[words[i][j+1:]]; has {
					ans = append(ans, words[i])
					break
				}
			}
		}
		rem[words[i]] = struct{}{}
	}
	return ans
}

func main() {
	maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	})
}
