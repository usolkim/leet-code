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
	matrix[0] = []int{5,4,4,3,9,2,6,0,9,0}
	matrix[1] = []int{5,4,4,3,1,5,2,6,5,6}
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

	row, col := 0, 0

	mp := 1<<uint(m+2) - 2
	np := 1<<uint(n+2) - 2

	for i := 0; i < m; i++ {
		row = (row << 1) + 1
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row &= np
			}
		}
	}

	for j := 0; j < n; j++ {
		col = (col << 1) + 1
		for i := 0; i < m; i++ {
			if matrix[i][j] == 0 {
				col &= mp
			}
		}
	}

	for i := 0; i < m; i++ {
		rtmp := 1 << uint(m-i-1)
		for j := 0; j < n; j++ {
			ctmp := 1 << uint(n-j-1)
			if row&rtmp == 0 || col&ctmp == 0 {
				matrix[i][j] = 0
			}
		}
	}
}
