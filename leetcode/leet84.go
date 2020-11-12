//
// Author: cfr
//

package main

func top(stack []int) int {
	return stack[len(stack)-1]
}
func largestRectangleArea(heights []int) int {
	var stack []int
	var res, with int
	res = 0
	stack = append(stack, -1)
	for i := 0; i < len(heights); i++ {
		for top(stack) != -1 && heights[top(stack)] < heights[i] {
			p := top(stack)
			stack = append(stack[0 : len(stack)-1])
			if top(stack) == -1 {
				with = i - p
			} else {
				with = i - top(stack) - 1
			}
			res = max(res, with*heights[p])
		}
		stack = append(stack, i)
	}

	for top(stack) != -1 {
		p := top(stack)
		stack = append(stack[0 : len(stack)-1])
		if top(stack) == -1 {
			with = len(heights) - p +1
		} else {
			with = len(heights) - top(stack)
		}
		res = max(res, with*heights[p])
	}
	return res
}

func max(res int, i int) int {
	if res > i {
		return res
	}
	return i
}
