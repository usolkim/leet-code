package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(5, 3))
}

func combine(n int, k int) [][]int {
	if k == 0 {
		return [][]int{[]int{}}
	} else if k == 1 {
		res := make([][]int, 0)
		for i := 1; i <= n; i++ {
			res = append(res, []int{i})
		}
		return res
	}
	res := make([][]int, 0)
	c := combine(n, k-1)
	for i := 1; i <= n-k+1; i++ {
		for _, v := range c {
			if v[0] > i {
				r := make([]int, 0)
				r = append(r, i)
				r = append(r, v...)
				res = append(res, r)
			}
		}
	}
	return res
}
