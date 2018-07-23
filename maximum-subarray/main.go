package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, -1, -3, 0, 0}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 14}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -5, 3, 12, -1, 3, 6}))
}

func maxSubArray(nums []int) int {
	l := len(nums)

	max := nums[0]
	s := 0

	for i := 0; i < l; i++ {
		s += nums[i]
		if s < nums[i] {
			s = nums[i]
		}
		if max < s {
			max = s
		}
	}

	return max
}
