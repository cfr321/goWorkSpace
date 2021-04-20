package main

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return nums[l]
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}
	return true
}
