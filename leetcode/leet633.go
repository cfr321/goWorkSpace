package main

import "math"

func judgeSquareSum(c int) bool {
	for i := 0; i*i <= c/2; i++ {
		tmp := int(math.Sqrt(float64(c - i*i)))
		if tmp*tmp == (c - i*i) {
			return true
		}
	}
	return false
}

type MapSum struct {
	sum  int
	sons [26]*MapSum
	rem  map[string]int
}

/** Initialize your data structure here. */
func Constructor3() MapSum {
	return MapSum{0, [26]*MapSum{}, make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	add := val - this.rem[key]
	this.rem[key] = val
	tmp := this
	this.sum += add
	for i := 0; i < len(key); i++ {
		pos := key[i] - 'a'
		if tmp.sons[pos] == nil {
			tmp.sons[pos] = &MapSum{0, [26]*MapSum{}, make(map[string]int)}
		}
		tmp.sons[pos].sum += add
		tmp = tmp.sons[pos]
	}
}

func (this *MapSum) Sum(prefix string) int {
	tmp := this
	for i := 0; i < len(prefix); i++ {
		if tmp.sons[prefix[i]-'a'] == nil {
			return 0
		}
		tmp = tmp.sons[prefix[i]-'a']
	}
	return tmp.sum
}
