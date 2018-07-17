package main

import "fmt"

func main() {
	fmt.Println(jump([]int{}))
	fmt.Println(jump([]int{1}))
	fmt.Println(jump([]int{1, 1}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{3, 1, 1, 2, 1, 4, 6, 6, 6}))
	fmt.Println(jump([]int{3, 1, 5, 2, 1, 4, 6, 6, 6}))
	fmt.Println(jump([]int{10, 1, 1, 2, 1, 4, 6, 6, 6}))
}

func jump(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}

	s := 0
	max := nums[0]
	jump := 1

	for s+max < l-1 {
		max = -1
		for j := 1; j <= nums[s]; j++ {
			if max < nums[s+j] {
				max = nums[s+j]
			}
		}
		s = s + max
		jump++
	}

	return jump
}
