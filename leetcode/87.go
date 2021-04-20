package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	l := len(s1)
	if len(s2) != l {
		return false
	}
	var dp [31][31][31]bool

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if s1[i] == s2[j] {
				dp[i][j][1] = true
			}
		}
	}
	for Len := 2; Len <= l; Len++ {
		for i := 0; i <= l-Len; i++ {
			for j := 0; j <= l-Len; j++ {
				for k := 1; k < Len; k++ {
					var son1, son2 bool
					son1 = dp[i][j][k] && dp[i+k][j+k][Len-k]
					son2 = dp[i+k][j][Len-k] && dp[i][j+Len-k][k]
					dp[i][j][Len] = son1 || son2
					if dp[i][j][Len] {
						break
					}
				}
			}
		}
	}
	return dp[0][0][l]
}

func main() {

	fmt.Println(isScramble("abcdbdacbdac", "bdacabcdbdac"))
}
