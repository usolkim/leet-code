package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 0))
	fmt.Println(combine(4, 1))
	fmt.Println(combine(4, 2))
	fmt.Println(combine(5, 3))
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	tmp := make([]int, k)

	var backt func(int, int)
	backt = func(i, v int) {
		if i == k+1 {
			tmp2 := make([]int, k)
			copy(tmp2, tmp)
			res = append(res, tmp2)
			return
		}

		for j := v; j <= n-k+i; j++ {
			tmp[i-1] = j
			backt(i+1, j+1)
		}
	}

	backt(1, 1)
	return res
}
