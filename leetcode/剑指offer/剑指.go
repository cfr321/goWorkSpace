package 剑指offer

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
