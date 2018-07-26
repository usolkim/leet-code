package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(0))
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(6))
	fmt.Println(climbStairs(7))
	fmt.Println(climbStairs(8))
	fmt.Println(climbStairs(9))
	fmt.Println(climbStairs(10))
	fmt.Println(climbStairs(4617213))
}

func climbStairs(n int) int {
	if n < 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	ar := make([]int, n)
	ar[0] = 1
	ar[1] = 2

	for i := 2; i < n; i++ {
		ar[i] = ar[i-2] + ar[i-1]
	}
	return ar[n-1]
}
