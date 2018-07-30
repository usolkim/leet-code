package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{1, 5, 3, 7}))
}

func subsets(nums []int) [][]int {
	l := len(nums)

	res := make([][]int, 0)
	tmp := make([]int, l)
	var backt func(int, int, int)

	backt = func(i, d, k int) {
		if i == k+1 {
			tmp2 := make([]int, k)
			copy(tmp2, tmp)
			res = append(res, tmp2)
			return
		}
		for j := d; j <= l-k+i; j++ {
			tmp[i-1] = nums[j-1]
			backt(i+1, j+1, k)
		}
	}

	for i := 0; i <= l; i++ {
		backt(1, 1, i)
	}

	return res
}
