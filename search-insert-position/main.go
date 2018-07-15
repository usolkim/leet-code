package main

import "fmt"

func main() {
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{}, 0))
	fmt.Println(searchInsert([]int{0}, 0))
	fmt.Println(searchInsert([]int{3}, 1))
	fmt.Println(searchInsert([]int{3}, 5))
}

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	s := 0
	e := l - 1
	i := 0

	for s <= e {
		i = (s + e) / 2
		if nums[i] < target {
			s = i + 1
		} else if nums[i] > target {
			e = i - 1
		} else {
			break
		}
	}

	if nums[i] < target {
		return i + 1
	}

	return i
}
