//
// Author: cfr
//

package main

func plusOne(digits []int) []int {
	l := len(digits)
	if l == 1 && digits[0] == 0 {
		digits[0]=1
		return digits
	}
	for i := l-1; i >=0 ; i-- {
		if digits[i]<9 {
			digits[i]++
			break
		}else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1},digits...)
	}
	return digits
}