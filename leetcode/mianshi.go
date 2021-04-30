package main

import (
	"sort"
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
		rem[s[i]] ++
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
		}else{
			if flag == 1 {
				return false
			}
			flag++
			if lena == lenb {
				a++
				b++
			}else if lena > lenb{
				a++
			}else {
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
