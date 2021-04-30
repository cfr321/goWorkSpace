package main

func canCross(stones []int) bool {
	if len(stones) <= 1 {
		return true
	}
	if stones[1]-stones[0] != 1 {
		return false
	}
	dp := make([][]int, len(stones))
	for i := 0; i < len(stones); i++ {
		dp[i] = make([]int,len(stones)+2)
	}
	dp[1][1] = 1
	for i := 2; i < len(stones); i++ {
		for j := 1; j < i; j++ {
			l := stones[i] - stones[j]
			if l < len(stones) {
				for k := l-1; k <=l+1 ; k++ {
					dp[i][l] |= dp[j][k]
				}
			}
		}
	}
	for i := 1; i < len(dp[0]) ; i++ {
		if dp[len(stones)-1][i] == 1 {
			return true
		}
	}
	return false
}
func main() {
	canCross([]int{0, 1, 3, 5, 6, 8, 12, 17})
}
