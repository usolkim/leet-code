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

	return calc(candidates, m, target)
}

func calc(c []int, m map[int]int, target int) [][]int {
	l := len(c)
	ret := make([][]int, 0)

	k := c[0]
	v := m[k]
	for i := 0; i <= v; i++ {
		e := k * i
		if e == target {
			tmp := make([]int, 0)
			for j := 0; j < i; j++ {
				tmp = append(tmp, k)
			}
			ret = append(ret, tmp)
		} else if e < target && v < l {
			r := calc(c[v:l], m, target-e)
			for _, vj := range r {
				tmp := make([]int, 0)
				for j := 0; j < i; j++ {
					tmp = append(tmp, k)
				}
				tmp = append(tmp, vj...)
				ret = append(ret, tmp)
			}
		}
	}

	return ret
}
