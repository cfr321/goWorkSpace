package main

import (
	"sort"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 01
func divide(a int, b int) int {
	if a == -(1<<31) && b == -1 {
		return 1<<31 - 1
	}
	flag := a^b >= 0
	a = abs(a)
	b = abs(b)
	res := 0
	for a >= b {
		shirt := 1
		for (b << shirt) <= a {
			shirt++
		}
		a -= b << (shirt - 1)
		res += 1 << (shirt - 1)
	}
	if flag {
		return res
	}
	return -res
}

// 02
func addBinary(a string, b string) string {
	la, lb := len(a)-1, len(b)-1
	lm := min(la, lb)
	be := byte(0)
	var res string
	for ; lm >= 0; lm-- {
		tmp := a[la] + b[lb] + be - 2*'0'
		be = tmp / 2
		tmp %= 2
		res = string(tmp) + res
		la--
		lb--
	}

	for ; la >= 0; la-- {
		tmp := a[la] + be - '0'
		be = tmp / 2
		tmp %= 2
		res = string(tmp) + res
	}

	for ; lb >= 0; lb-- {
		tmp := b[lb] + be - '0'
		be = tmp / 2
		tmp %= 2
		res = string(tmp) + res
	}
	if be == 1 {
		res = string('1') + res
	}
	return res
}

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < 64; i++ {
		tmp := 0
		for _, num := range nums {
			if num&(1<<i) != 0 {
				tmp++
			}
		}
		if tmp%3 != 0 {
			res |= 1 << i
		}
	}
	return res
}

// 05
func maxProduct(words []string) int {
	rem := make([]int, len(words))
	for j, word := range words {
		for i := 0; i < len(word); i++ {
			rem[j] |= 1 << (word[i] - 'a')
		}
	}
	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if (rem[i] & rem[j]) == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
		}
	}
	return res
}

// 06
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for tmp := numbers[l] + numbers[r]; tmp != target; {
		if tmp > target {
			r--
		} else {
			l++
		}
	}
	return []int{l, r}
}

// 07 三数之和，定1求2
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			tmp := nums[l] + nums[r] + nums[i]
			if tmp == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if tmp > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

// 08 最短子数组
func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := 0
	res := len(nums) + 1
	for r < len(nums) {
		sum += nums[r]
		r++
		for sum >= target {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// 09
func numSubarrayProductLessThanK(nums []int, k int) int {
	l := 0
	res := 0
	tmp := 1
	for r, num := range nums {
		tmp *= num
		for tmp >= k {
			tmp /= nums[l]
			l++
		}
		res += r - l + 1
	}
	return res
}

// 10
func subarraySum(nums []int, k int) int {
	rem := make(map[int]int)
	rem[0] = 1
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res += rem[sum-k]
		rem[sum]++
	}
	return res
}

// 11
func findMaxLength(nums []int) int {
	tmp := 0
	res := 0
	rem := make(map[int]int)
	rem[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
		} else {
			tmp--
		}
		if k, has := rem[tmp]; has {
			res = max(res, i-k)
		} else {
			rem[tmp] = i
		}
	}
	return res
}

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}
