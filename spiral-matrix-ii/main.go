package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(0))
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
	fmt.Println(generateMatrix(5))
	fmt.Println(generateMatrix(6))
	fmt.Println(generateMatrix(7))
}

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	} else if n == 1 {
		return [][]int{[]int{1}}
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	l := n * n
	a, j, k, r := 0, 0, 0, (n-1)*4

	for i := 1; i <= l; i++ {
		if i == r {
			a++
			r = ((a+1) * n - ((a+1)*(a+1))) * 4
		}
		res[j][k] = i
		if j == a && k < n-a-1 {
			k++
		} else if k == n-a-1 && j < n-a-1 {
			j++
		} else if j == n-a-1 && k > a {
			k--
		} else {
			j--
		}
	}

	return res
}
