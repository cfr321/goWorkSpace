package main

func numRabbits(answers []int) int {
	m := make(map[int]int)
	for i := 0; i < len(answers); i++ {
		m[answers[i]]++
	}
	res := 0
	for i := range m {
		res += (m[i] + i) / (i + 1)
	}
	return res
}
