package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{1, 1, 1, 1, 1, 2, 2, 2, 2}, 5))
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

func combinationSum2(candidates []int, target int) [][]int {
	if target < 1 {
		return nil
	}
	l := len(candidates)
	if l == 0 {
		return nil
	}
	if l == 1 && candidates[0] != target {
		return nil
	}

	sort.Ints(candidates)

	m := make(map[int]int, 0)

	for _, v := range candidates {
		m[v]++
	}

	ret := make([][]int, 0)
	r := make([]int, 0)

	calc(&ret, candidates, r, m, target)

	return ret
}

func calc(ret *[][]int, c []int, r []int, m map[int]int, target int) {
	l := len(c)

	k := c[0]
	v := m[k]
	for i := 0; i <= v; i++ {
		e := k * i
		if e == target {
			r = append(r, c[:i]...)
			*ret = append(*ret, r)
			return
		} else if e < target && v < l && c[v] <= target-e {
			tmp := make([]int, len(r))
			copy(tmp, r)
			tmp = append(tmp, c[:i]...)
			calc(ret, c[v:], tmp, m, target-e)
		}
	}
}
