package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{1}))
	fmt.Println(canJump([]int{2, 3, 1, 0, 2, 1, 1, 1, 1, 1, 1}))
}

func canJump(nums []int) bool {
	l := len(nums)
	if l <= 1 {
		return true
	}

	j := nums[0]

	if j == 0 {
		return false
	}
	s := 1
	m := j
	for j < l-1 {
		for i := s; i <= j; i++ {
			if m < i+nums[i] {
				m = i + nums[i]
			}
		}
		s = j + 1
		j = m
		if s > j {
			return false
		}
	}
	return true
}
