//
// Author: cfr
//

package main
func grayCode(n int) []int {
	var res =[]int{0}
	if n==0 {
		return res
	}
	res = append(res, 1)
	var bit = 1
	for i := 1; i < n ; i++ {
		bit<<=1
		for j := len(res)-1; j >=0 ; j-- {
			res = append(res, res[j]|bit)
		}
	}
	return res
}