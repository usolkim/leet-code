package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 19, 25, 38, 43, 57, 66, 71, 84}
	nums2 := []int{3, 4, 13, 28, 33, 37, 49, 51, 52, 63, 74, 83, 86, 91, 94, 99}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2, 19, 25, 38, 43, 57, 66, 71, 84}
	nums2 = []int{3, 4, 13, 28, 33, 37, 49, 51, 52, 63, 74, 83, 86, 91, 94}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 3}
	nums2 = []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 > l2 {
		nt := nums1
		nums2 = nums1
		nums1 = nt
		lt := l1
		l1 = l2
		l2 = lt
	}

	var s, e, m2d int = 0, l1, int((l1 + l2 + 1) / 2)
	var isOdd bool
	var res float64
	if (l1+l2)%2 == 0 {
		isOdd = true
	}

	for i := 0; i < 10; i++ {
		m1 := int((s + e) / 2)
		m2 := m2d - m1

		fmt.Println(m1, m2)

		if nums2[m2-1] > nums1[m1] {
			if m1 + 1 == l1 {
				s = l1
			} else {
				s = m1
			}
		} else if nums1[m1-1] > nums2[m2] {
			e = m1 + 1
		} else {
			if nums1[m1-1] < nums2[m2-1] {
				res = float64(nums2[m2-1])
			} else {
				res = float64(nums1[m1-1])
			}
			if isOdd {
				if nums1[m1] < nums2[m2] {
					res = (res + float64(nums1[m1])) / 2
				} else {
					res = (res + float64(nums2[m2])) / 2
				}
			}
			break
		}
	}

	return res
}
