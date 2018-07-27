package main

import (
	"fmt"
)

func main() {
	matrix := make([][]int, 3)
	matrix[0] = []int{1, 1, 1}
	matrix[1] = []int{1, 0, 1}
	matrix[2] = []int{1, 1, 1}
	print(matrix)
	setZeroes(matrix)
	print(matrix)

	matrix = make([][]int, 3)
	matrix[0] = []int{0, 1, 2, 0}
	matrix[1] = []int{3, 4, 5, 2}
	matrix[2] = []int{1, 3, 1, 5}
	print(matrix)
	setZeroes(matrix)
	print(matrix)

	matrix = make([][]int, 3)
	matrix[0] = []int{-4, -2147483648, 6, -7, 0}
	matrix[1] = []int{-8, 6, -8, -6, 0}
	matrix[2] = []int{2147483647, 2, -9, -6, -10}
	print(matrix)
	setZeroes(matrix)
	print(matrix)

	matrix = make([][]int, 2)
	matrix[0] = []int{5, 4, 4, 3, 9, 2, 6, 0, 9, 0}
	matrix[1] = []int{5, 4, 4, 3, 1, 5, 2, 6, 5, 6}
	print(matrix)
	setZeroes(matrix)
	print(matrix)

	matrix = make([][]int, 5)
	matrix[0] = []int{8, 3, 1, 4, 6, 4, 0, 3, 4}
	matrix[1] = []int{9, 1, 3, 0, 1, 5, 7, 4, 1}
	matrix[2] = []int{2, 2, 5, 2147483647, 5, 4, 4, 3, 8}
	matrix[3] = []int{4, 9, 7, 0, 3, 6, 9, 5, 9}
	matrix[4] = []int{4, 1, 8, 8, 4, 1, 5, 7, 6}
	print(matrix)
	setZeroes(matrix)
	print(matrix)

	matrix = make([][]int, 8)
	matrix[0] = []int{3, 5, 5, 6, 9, 1, 4, 5, 0, 5}
	matrix[1] = []int{2, 7, 9, 5, 9, 5, 4, 9, 6, 8}
	matrix[2] = []int{6, 0, 7, 8, 1, 0, 1, 6, 8, 1}
	matrix[3] = []int{7, 2, 6, 5, 8, 5, 6, 5, 0, 6}
	matrix[4] = []int{2, 3, 3, 1, 0, 4, 6, 5, 3, 5}
	matrix[5] = []int{5, 9, 7, 3, 8, 8, 5, 1, 4, 3}
	matrix[6] = []int{2, 4, 7, 9, 9, 8, 4, 7, 3, 7}
	matrix[7] = []int{3, 5, 2, 8, 8, 2, 2, 4, 9, 8}
	print(matrix)
	setZeroes(matrix)
	print(matrix)

}

func print(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("-------------------")
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
	if n == 0 {
		return
	}

	frow := false
	fcol := false

	if matrix[0][0] == 0 {
		fcol = true
		frow = true
	}

	if !fcol {
		for i := 1; i < m; i++ {
			if matrix[i][0] == 0 {
				fcol = true
				break
			}
		}
	}

	if !frow {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 {
				frow = true
				break
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// print(matrix)

	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	if fcol {
		matrix[0][0] = 0
		for i := 1; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if frow {
		matrix[0][0] = 0
		for j := 1; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}
