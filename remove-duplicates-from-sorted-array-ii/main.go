package main

import (
	"fmt"
)

func main() {
	var l int
	var nums []int

	nums = []int{1, 1, 1, 2, 2, 3}
	l = removeDuplicates(nums)
	print(nums, l)

	nums = []int{1, 1, 1, 1, 2, 2, 2, 3}
	l = removeDuplicates(nums)
	print(nums, l)

	nums = []int{}
	l = removeDuplicates(nums)
	print(nums, l)

	nums = []int{1}
	l = removeDuplicates(nums)
	print(nums, l)

	nums = []int{1, 1}
	l = removeDuplicates(nums)
	print(nums, l)

	nums = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	l = removeDuplicates(nums)
	print(nums, l)

	nums = []int{1, 3}
	l = removeDuplicates(nums)
	print(nums, l)

	nums = []int{1, 3, 5}
	l = removeDuplicates(nums)
	print(nums, l)

	nums = []int{0, 0, 0, 1, 3, 3, 3, 7, 8, 8, 8}
	l = removeDuplicates(nums)
	print(nums, l)

	nums = []int{0, 0, 0, 1, 3, 3, 3, 7, 8, 8, 9}
	l = removeDuplicates(nums)
	print(nums, l)
}

func print(nums []int, l int) {
	fmt.Print("\"")
	for i := 0; i < l; i++ {
		fmt.Print(nums[i], ",")
	}
	fmt.Println("\"")
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 3 {
		return l
	}
	start, end := 2, 2

	for end < l {
		if nums[start-2] >= nums[start] {
			for end < l-1 && nums[start-2] >= nums[end] {
				end++
			}
			if nums[start-2] < nums[end] {
				nums[start], nums[end] = nums[end], nums[start]
				start++
			}
		} else {
			start++
		}
		end++
	}
	return start
}
