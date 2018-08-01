package main

import (
	"fmt"
)

func main() {
	// fmt.Println(search([]int{3, 3}, 3))
	fmt.Println(search([]int{1, 3, 1, 1, 1}, 3))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{0, 1, 2, 2, 5, 6, 0}, 0))
	fmt.Println(search([]int{0, 0, 1, 2, 2, 5, 6}, 0))
	fmt.Println(search([]int{6, 0, 0, 1, 2, 2, 5}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 1, 2}, 0))
	fmt.Println(search([]int{0, 1, 1, 2, 2, 5, 6, 0}, 0))
	fmt.Println(search([]int{0, 0, 1, 1, 2, 2, 5, 6}, 0))
	fmt.Println(search([]int{6, 0, 0, 1, 1, 2, 2, 5}, 0))
	fmt.Println(search([]int{2, 4, 4, 4, 5, 6, 0, 0, 1, 1, 2}, 0))
	fmt.Println(search([]int{2, 4, 4, 4, 5, 6, 0, 0, 1, 1, 2}, 0))
	fmt.Println(search([]int{0, 1, 1, 2, 2, 4, 4, 4, 5, 6, 0}, 0))
	fmt.Println(search([]int{0, 0, 1, 1, 2, 2, 4, 4, 4, 5, 6}, 0))
	fmt.Println(search([]int{6, 0, 0, 1, 1, 2, 2, 4, 4, 4, 5}, 0))
	fmt.Println(search([]int{3}, 3))
	fmt.Println(search([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0))
	fmt.Println("==================")
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{0, 1, 2, 2, 5, 6, 0}, 3))
	fmt.Println(search([]int{0, 0, 1, 2, 2, 5, 6}, 3))
	fmt.Println(search([]int{6, 0, 0, 1, 2, 2, 5}, 3))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 1, 2}, 3))
	fmt.Println(search([]int{0, 1, 1, 2, 2, 5, 6, 0}, 3))
	fmt.Println(search([]int{0, 0, 1, 1, 2, 2, 5, 6}, 3))
	fmt.Println(search([]int{6, 0, 0, 1, 1, 2, 2, 5}, 3))
	fmt.Println(search([]int{2, 4, 4, 4, 5, 6, 0, 0, 1, 1, 2}, 3))
	fmt.Println(search([]int{0, 1, 1, 2, 2, 4, 4, 4, 5, 6, 0}, 3))
	fmt.Println(search([]int{0, 0, 1, 1, 2, 2, 4, 4, 4, 5, 6}, 3))
	fmt.Println(search([]int{6, 0, 0, 1, 1, 2, 2, 4, 4, 4, 5}, 3))
	fmt.Println(search([]int{}, 3))
	fmt.Println(search([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 3))
	fmt.Println(search([]int{2, 4, 4, 4, 5, 6, 0, 0, 1, 1, 2}, -1))
	fmt.Println(search([]int{2, 4, 4, 4, 5, 6, 0, 0, 1, 1, 2}, 8))
}

func search(nums []int, target int) bool {
	l := len(nums)
	if l == 0 {
		return false
	} else if l == 1 {
		return nums[0] == target
	}
	first := 0
	if nums[0] >= nums[l-1] {
		for i := 1; i < l; i++ {
			if nums[i-1] > nums[i] {
				first = i
				break
			}
		}
	}
	if first == 0 {
		return bsearch(nums, 0, l, target)
	} else if nums[0] <= target && target <= nums[first-1] {
		return bsearch(nums, 0, first, target)
	} else if nums[first] <= target && target <= nums[l-1] {
		return bsearch(nums, first, l, target)
	}
	return false
}

func bsearch(nums []int, s, e, target int) bool {
	if s == e {
		return len(nums) > s && nums[s] == target
	}
	c := (s + e) / 2
	if nums[c] == target {
		return true
	} else if nums[c] < target {
		return bsearch(nums, c+1, e, target)
	}
	return bsearch(nums, s, c, target)
}
