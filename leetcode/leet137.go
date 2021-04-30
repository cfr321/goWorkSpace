package main

func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		tmp := 0
		for j := 0; j < len(nums); j++ {
			if (nums[j]>>i)&1 != 0 {
				tmp++
			}
		}
		if tmp%3 != 0 {
			ans |= 1<<i
		}
	}
	return ans
}
