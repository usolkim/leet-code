package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{1}
	sortColors(nums)
	fmt.Println(nums)

	nums = []int{1, 1, 1, 2, 2, 2, 0, 0, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}
	min, max := 0, l-1

	for i := 0; i <= max; i++ {
		if nums[i] == 0 {
			nums[min], nums[i] = nums[i], nums[min]
			min++
		} else if nums[i] == 2 {
			nums[max], nums[i] = nums[i], nums[max]
			i--
			max--
		}
	}
}
