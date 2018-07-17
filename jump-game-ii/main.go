package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{2}))
	fmt.Println(jump([]int{1, 1}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{3, 1, 5, 1, 2, 3, 1, 4, 2}))
	fmt.Println(jump([]int{30, 1, 1, 1, 1, 1, 1, 5, 1, 2, 3, 1, 4, 2}))
}

func jump(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}

	s := 1
	max := nums[0]
	count := 1

	for max < l-1 {
		omax := max
		max = 0
		for i := s; i <= omax; i++ {
			if max < i+nums[i] {
				max = i + nums[i]
			}
		}
		s = omax + 1
		count++
	}

	return count
}
