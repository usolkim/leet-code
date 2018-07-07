package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{-1, 2, -1, 1, 0, 3, -2}, 0))
	fmt.Println(fourSum([]int{5, 5, 3, 5, 1, -5, 1, -2},
		4))
}

func fourSum(nums []int, target int) [][]int {
	ln := len(nums)
	if ln < 4 {
		return nil
	}
	ret := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < ln-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := ln - 1; j > i+2; j-- {
			if j < ln-1 && nums[j] == nums[j+1] {
				continue
			}
			for k, l := i+1, j-1; k < l; {
				s := nums[i] + nums[j] + nums[k] + nums[l]
				if s == target {
					ret = append(ret, []int{nums[i], nums[k], nums[l], nums[j]})
					k++
					l--
					for k < l && nums[k-1] == nums[k] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if s < target {
					k++
				} else {
					l--
				}
			}

		}
	}
	return ret
}
