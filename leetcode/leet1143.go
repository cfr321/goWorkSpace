package main

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)
	var tmp, pre int

	for i := 0; i < len(text1); i++ {
		tmp = 0
		for j := 0; j < len(text2); j++ {

			//  dp[j+1]在没有修改之前就是dp[i][j+1]
			//  当前求解是 dp[i+1][j+1] =  （不相等） max(dp[i][j+1],dp[i+1][j])  or (相等） dp[i][j] + 1

			pre = tmp
			tmp = dp[j+1]
			if text2[i] == text1[j] {
				dp[j+1] = pre + 1
			} else {
				dp[j+1] = max(dp[j+1], dp[j-1])
			}
		}
	}
	return dp[len(text2)]
}

func shipWithinDays(weights []int, D int) int {
	sum := weights[0]
	for i := 1; i < len(weights); i++ {
		sum += weights[i]
	}
	res := (sum + D - 1) / D
	for true {
		k := 0
		for i := 0; i < D && k < len(weights); i++ {
			sum := 0
			for k < len(weights) {
				if (sum + weights[k]) <= res {
					sum += weights[k]
					k++
				} else {
					break
				}
			}
		}
		if k == len(weights) {
			break
		}
	}
	return res
}
