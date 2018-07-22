package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 1, 2, 3}))
	fmt.Println(permuteUnique([]int{1, 1, 2, 3, 1}))
	fmt.Println(permuteUnique([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(permuteUnique([]int{1, 3, 2, 3, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 3, 2, 3, 1, 2, 1}))
	fmt.Println(permuteUnique([]int{1, 1, 1}))
	fmt.Println(permuteUnique([]int{1}))
	fmt.Println(permuteUnique([]int{}))

}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	permute(nums, 0, &res)
	return res
}

func permute(nums []int, start int, res *[][]int) {
	if start >= len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		check := true
		for j := start; j < i; j++ {
			if nums[j] == nums[i] {
				check = false
				break
			}
		}
		if check {
			nums[start], nums[i] = nums[i], nums[start]
			permute(nums, start+1, res)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

}
