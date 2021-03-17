//
// Author: cfr
//

package main

import (
	"math/rand"
)

type RandomizedCollection struct {
	m    map[int][]int
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	this := RandomizedCollection{m: make(map[int][]int)}
	return this
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.nums = append(this.nums, val)
	res,has := this.m[val]
	if !has {
		res = make([]int,0)
	}
	res = append(res, len(this.nums)-1)
	return !has
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	ids,has:=this.m[val]
	if !has {
		return false
	}
	lids:=len(ids)
	n:=len(this.nums)
	i:=ids[lids-1]
	this.nums[i] = this.nums[n-1]         //最后一个数前调
	ids=ids[:lids-1]              //删掉val对应数组的最后以为
	this.m[this.nums[n-1]] = this.m[this.nums[n-1]][0:len(this.m[this.nums[n-1]])-1]
	if i < n-1 {
		this.m[this.nums[n-1]] = append(this.m[this.nums[n-1]], i)
	}
	if len(this.m[val]) == 0 {
		delete(this.m,val)
	}
	this.nums = this.nums[:n-1]
	return true
}

//func equalSubstring(s string, t string, maxCost int) int {
//	rem:= make([]int, len(s))
//
//	for i := 0; i < len(s); i++ {
//
//	}
//}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

