package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchRange([]int{1, 1, 2}, 1))
	fmt.Println(searchRange([]int{}, 8))
	fmt.Println(searchRange([]int{0}, 8))
	fmt.Println(searchRange([]int{8}, 8))
	fmt.Println(searchRange([]int{8, 8}, 8))
	fmt.Println(searchRange([]int{7, 8, 9}, 8))
	fmt.Println(searchRange([]int{1, 8, 9}, 8))
	fmt.Println(searchRange([]int{1, 8, 8, 9}, 8))
	fmt.Println(searchRange([]int{1, 1, 8, 9, 9}, 8))
	fmt.Println(searchRange([]int{1, 1, 8, 8, 9, 9}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	}
	if l == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
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
			if i-1 >= 0 && nums[i] == nums[i-1] {
				e = i - 1
				continue
			}
			break
		}
	}

	if nums[i] != target {
		return []int{-1, -1}
	}

	left := i

	s = i
	e = l - 1

	for s <= e {
		i = (s + e) / 2
		if nums[i] < target {
			s = i + 1
		} else if nums[i] > target {
			e = i - 1
		} else {
			if i+1 < l && nums[i] == nums[i+1] {
				s = i + 1
				continue
			}
			break
		}
	}

	right := i

	return []int{left, right}
}
