package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{3, 2, 1}))
	fmt.Println(permute([]int{1, 2, 3, 4}))
	fmt.Println(permute([]int{1, 4, 2, 3}))
	fmt.Println(permute([]int{1}))
	fmt.Println(permute([]int{}))
}

func permute(nums []int) [][]int {
	l := len(nums)
	if l <= 1 {
		return [][]int{nums}
	} else if l == 2 {
		return [][]int{nums, []int{nums[1], nums[0]}}
	}

	bf := permute(nums[1:])

	res := make([][]int, 0)
	n := nums[0]
	for _, b := range bf {
		for i := 0; i < l; i++ {
			tmp := make([]int, 0)
			tmp = append(tmp, b[:i]...)
			tmp = append(tmp, n)
			tmp = append(tmp, b[i:]...)
			res = append(res, tmp)
		}
	}

	return res

}
