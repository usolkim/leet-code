package main

import (
	"fmt"
)

func main() {
	fmt.Println(grayCode(0))
	fmt.Println(grayCode(1))
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(3))
	fmt.Println(grayCode(4))
	fmt.Println(grayCode(5))
	fmt.Println(grayCode(6))
}

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := make([]int, 1<<uint(n))
	res[0] = 0
	res[1] = 1
	for i := 2; i <= n; i++ {
		s := 1 << uint(i-1)
		for j := 0; j < s; j++ {
			res[s+j] = s + res[s-j-1]
		}
	}
	return res
}
