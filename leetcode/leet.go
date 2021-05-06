package main

import (
	"math"
	"sort"
)

//Definition for Employee.
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
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
func decode(encoded []int, first int) []int {
	var tmp int
	for i := 0; i < len(encoded); i++ {
		tmp = first
		first = encoded[i] ^ first
		encoded[i] = tmp
	}
	encoded = append(encoded, first)
	return encoded
}
