//
// Author: cfr
//

package myUtil

func Swap(a, b *int) {
	tmp := a
	a = b
	b = tmp
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
