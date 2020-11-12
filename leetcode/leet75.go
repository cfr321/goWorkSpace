//
// Author: cfr
//

package main

func swap75(nums []int, l, r int) {
	nums[l], nums[r] = nums[r], nums[l]
}
func sortColors(nums []int) {
	l, r := -1, len(nums)
	k := 0
	for k < r {
		if nums[k] == 1 {
			k++
		} else if nums[k]== 2 {
			r--
			swap75(nums,k,r)
		} else {
			l++
			swap75(nums,l,k)
			k++
		}
	}
}
