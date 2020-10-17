//
// Author: cfr
//

package main

import (
	"fmt"
	"strconv"
)

func swap(nums []int, l, r int) {
	temp:=nums[l]
	nums[l]=nums[r]
	nums[r]=temp
}
func reverce(nums []int,l,r int) {
	len:=r-l+1
	for i := 0; i < len/2; i++ {
		swap(nums,l+i,r-i)
	}
}
func nextPermutation(nums []int) {
	var i int
	for i = len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			var pos int
			for j := i; j < len(nums); j++ {
				if nums[j] > nums[i-1] && (j+1>=len(nums)|| nums[j+1]<=nums[i-1]) {
					pos = j
					break
				}
			}
			swap(nums,i-1,pos)
			reverce(nums,i,len(nums)-1)
			break
		}
	}
	if i == 0 {
		reverce(nums,0,len(nums)-1)
	}
}
func getPermutation(n int, k int) string {
	m := 1
	t := 1
	for m <= k {
		t++
		m *= t
	}
	var buff []int
	for i := 1; i <= n; i++ {
		buff = append(buff,i)
	}
	reverce(buff,n-t+1,n-1)
	m/=t
	k-=m
	for k > 0 {
		nextPermutation(buff)
		k--
	}
	s:=""
	for i := 0; i < len(buff); i++ {
		s+=strconv.Itoa(buff[i])
	}
	return s
}

func main() {
	fmt.Print(getPermutation(3,3))
}
