package main

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
