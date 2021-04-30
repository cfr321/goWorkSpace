//
// Author: cfr
//

package main

import (
	"strconv"
	"strings"
)

//01.06
func compressString(S string) string {
	var res []byte
	var l = 1
	for i := 0; i < len(S); i++ {
		if i == len(S)-1 || S[i] != S[i+1] {
			res = append(res, S[i])
			res = append(res, strconv.Itoa(l)...)
			l = 1
		} else {
			l++
		}
	}
	if len(res) >= len(S) {
		return S
	}
	return string(res)
}

//01.07
func swap1(a, b *int) {
	*a, *b = *b, *a
}
func rotate(matrix [][]int) {
	N := len(matrix)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			swap1(&matrix[i][j], &matrix[j][i])
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N/2; j++ {
			swap1(&matrix[i][j], &matrix[i][N-j-1])
		}
	}
}

//01.08
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

//01.09   !!!!!!
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	new := s2 + s2
	index := strings.Index(new, s1)
	return index != -1

	//for i := 0; i < len(s1); i++ {
	//	tmp:=s1[i:]+s1[0:i]
	//	if tmp == s2 {
	//		return true
	//	}
	//}
	//return false
}
