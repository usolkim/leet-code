package main

import (
	"fmt"
)

func main() {
	matrix := make([][]int, 0)
	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})
	rotate(matrix)

	for _, v := range matrix {
		fmt.Println(v)
	}

	matrix = make([][]int, 0)
	matrix = append(matrix, []int{5, 1, 9, 11})
	matrix = append(matrix, []int{2, 4, 8, 10})
	matrix = append(matrix, []int{13, 3, 6, 7})
	matrix = append(matrix, []int{15, 14, 12, 16})
	rotate(matrix)

	for _, v := range matrix {
		fmt.Println(v)
	}
}

func rotate(matrix [][]int) {
	l := len(matrix)

	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-j-1] = matrix[i][l-j-1], matrix[i][j]
		}
	}

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			matrix[i][j], matrix[l-j-1][l-i-1] = matrix[l-j-1][l-i-1], matrix[i][j]
		}
	}
}
