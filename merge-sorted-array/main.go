package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)

	nums1 = []int{1, 0, 0, 0, 0, 0}
	merge(nums1, 1, []int{2, 5, 6}, 3)
	fmt.Println(nums1)

	nums1 = []int{0, 0, 0, 0, 0}
	merge(nums1, 0, []int{2, 5, 6}, 3)
	fmt.Println(nums1)

	nums1 = []int{1, 2, 3, 0, 0}
	merge(nums1, 3, []int{}, 0)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1[:m+n])
}
