//
// Author: cfr
//

package main

func minWindow(s string, t string) string {
	m := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	var rl, rr = 0, len(s) + 1
	k := 0
	b := -1
	for i := 0; i < len(s); i++ {
		if m[s[i]] != 0 {
			if b == -1 {
				b = i
			}
			m2[s[i]]++
			if m2[s[i]] <= m[s[i]] {
				k++
			}
		}
		for k == len(t) {
			if rr-rl > i-b {
				rr = i
				rl = b
			}
			if m2[s[b]] > 0 {
				m2[s[b]]--
			}
			if m2[s[b]] < m[s[b]] {
				k--
			}
			b++
		}
	}
	if rr == len(s)+1 {
		return ""
	}
	return s[rl : rr+1]
}
