package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(1, 1))
	fmt.Println(uniquePaths(1, 100))
	fmt.Println(uniquePaths(100, 1))
	fmt.Println(uniquePaths(12, 17))
}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	ar := make([]int, m)

	ar[0] = 1

	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			ar[j] += ar[j-1]
		}
	}

	return ar[m-1]
}
