//
// Author: cfr
//

package main

import (
	"fmt"
	"sort"
)

func wordBreak(s string, wordDict []string) []string {
	m := make(map[string]int)
	for i := range wordDict {
		m[wordDict[i]] = 1
	}

	var dp1 [1000]bool

	dp1[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if m[s[j:i]] == 1 && dp1[j] {
				dp1[i] = true
			}
		}
	}

	var dp [1000][]string
	dp[0] = []string{""}
	if dp1[len(s)] {
		for i := 1; i <= len(s); i++ {
			for j := 0; j < i; j++ {
				if m[s[j:i]] == 1 {
					for k := range dp[j] {
						if dp[j][k] == "" {
							dp[i] = append(dp[i], s[0:i])
						} else {
							dp[i] = append(dp[i], dp[j][k]+" "+s[j:i])
						}
					}
				}
			}
		}
	}
	return dp[len(s)]
}
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var s []string
	s = append(s, beginWord)
	last := s[0]
	k := 2
	for len(s) != 0 {
		beginWord = s[0]
		s = s[1:]
		for i := range wordList {
			if wordList[i] != "" && diffecent(beginWord, wordList[i]) == 1 {
				if wordList[i] == endWord {
					return k
				}
				s = append(s, wordList[i])
				wordList[i] = ""
			}
		}
		if last == beginWord {
			k++
			if len(s) > 0 {
				last = s[len(s)-1]
			}
		}
	}
	return 0
}
func diffecent(word string, s string) int {
	var k = 0
	for i := 0; i < len(word); i++ {
		if word[i] != s[i] {
			k++
		}
	}
	return k
}

type bitarr struct {
	a  []int
	n  []int
}

func (b bitarr) Len() int {
	return len(b.a)
}

func (b bitarr) Less(i, j int) bool {
	if b.n[i] == b.n[j]{
		return b.a[i]<b.a[j]
	}
	return b.n[i] < b.n[j]
}

func (b bitarr) Swap(i, j int) {
	b.a[i],b.a[j] = b.a[j],b.a[i]
	b.n[i],b.n[j] = b.n[j],b.n[i]
}

func numBit(a int) int {
	i:=0
	for a != 0 {
		i++
		a &= a-1
	}
	return i
}
func sortByBits(arr []int) []int {
	var bitArr bitarr
	for i := 0; i < len(arr); i++ {
		bitArr.a = append(bitArr.a, arr[i])
		bitArr.n = append(bitArr.n, numBit(arr[i]))
	}
	sort.Sort(bitArr)
	return bitArr.a
}

func main() {
	fmt.Println(numBit(6978))
	fmt.Println(numBit(8641))
}
