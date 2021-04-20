package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}

	p1 := n
	p2 := 0
	i := 0
	for ; i < m+n && p1 < m+n && p2 < n; i++ {
		if nums1[p1] > nums2[p2] {
			nums1[i] = nums2[p2]
			p2++
		} else {
			nums1[i] = nums1[p1]
			p1++
		}
	}
	for p2 < n {
		nums1[i] = nums2[p2]
		i++
		p2++
	}

}
