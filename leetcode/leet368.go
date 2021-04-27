package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	var res [][]int
	sort.Ints(nums)
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		maxL := 0
		for j := 1; j < len(res); j++ {
			if nums[i]%res[j][len(res[j])-1] == 0 && len(res[j]) > len(res[maxL]) {
				maxL = j
			}
		}
		tmp := make([]int, len(res[maxL]))
		copy(tmp, res[maxL])
		res = append(res, tmp)
		res[len(res)-1] = append(res[len(res)-1], nums[i])
	}
	maxi := 0
	for i := 1; i < len(res); i++ {
		if len(res[maxi]) < len(res[i]) {
			maxi = i
		}
	}

	return res[maxi]
}
