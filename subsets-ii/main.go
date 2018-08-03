package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{}))
	fmt.Println(subsetsWithDup([]int{1}))
	fmt.Println(subsetsWithDup([]int{1, 1}))
	fmt.Println(subsetsWithDup([]int{1, 2}))
	fmt.Println(subsetsWithDup([]int{1, 2, 3}))
	fmt.Println(subsetsWithDup([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}))
}

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)

	n := []int{nums[0]}
	v := []int{1}
	l := 0

	for i := 1; i < len(nums); i++ {
		if n[l] < nums[i] {
			l++
			n = append(n, nums[i])
			v = append(v, 1)
		} else {
			v[l]++
		}
	}

	res := make([][]int, 0)

	var dfs func(int, []int)
	dfs = func(i int, t []int) {
		if i > l {
			tmp := make([]int, len(t))
			copy(tmp, t)
			res = append(res, tmp)
			return
		}
		for j := 0; j <= v[i]; j++ {
			dfs(i+1, t)
			t = append(t, n[i])
		}
	}

	dfs(0, []int{})

	return res
}
