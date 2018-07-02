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
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 {
		nums_t := nums1
		nums1 = nums2
		nums2 = nums_t
		len_t := len1
		len1 = len2
		len2 = len_t
	}

	var start, end, m2d int = 0, len1, int((len1 + len2 + 1) / 2)
	var isOdd bool
	var res float64
	if (len1+len2)%2 == 0 {
		isOdd = true
	}

	for start <= end {
		m1 := int((start + end) / 2)
		m2 := m2d - m1

		if m1 < len1 && nums2[m2-1] > nums1[m1] {
			start = m1 + 1
		} else if m1 > 0 && nums1[m1-1] > nums2[m2] {
			end = m1 - 1
		} else {
			if m1 == 0 {
				res = float64(nums2[m2-1])
			} else if m2 == 0 {
				res = float64(nums1[m1-1])
			} else if nums1[m1-1] < nums2[m2-1] {
				res = float64(nums2[m2-1])
			} else {
				res = float64(nums1[m1-1])
			}
			if isOdd {
				if m1 == len1 {
					res = (res + float64(nums2[m2])) / 2
				} else if m2 == len2 {
					res = (res + float64(nums1[m1])) / 2
				} else if nums1[m1] < nums2[m2] {
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
