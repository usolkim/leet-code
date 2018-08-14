package main

import (
	"fmt"
)

func main() {
	fmt.Println(numTrees(0))
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
	fmt.Println(numTrees(5))
	fmt.Println(numTrees(6))
	fmt.Println(numTrees(7))
	fmt.Println(numTrees(8))
	fmt.Println(numTrees(9))
}

func numTrees(n int) int {
	if n < 2 {
		return 1
	}

	a := make([]int, n+1)
	a[0] = 1
	a[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j <= (i-1)/2; j++ {
			k := i - 1 - j
			if j == k {
				a[i] += a[j] * a[j]
			} else {
				a[i] += 2 * a[j] * a[i-1-j]
			}
		}
	}

	return a[n]
}
