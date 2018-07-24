package main

import (
	"fmt"
)

func main() {
	fmt.Println(spiralOrder([][]int{[]int{3}, []int{2}}))
	fmt.Println(spiralOrder([][]int{}))
	fmt.Println(spiralOrder([][]int{[]int{1}}))
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3}}))
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	l := m * n

	res := make([]int, l)

	a, t, j, k := 0, 1, 0, 0

	for i := 0; i < l; i++ {
		res[i] = matrix[j][k]
		switch t {
		case 1:
			if k == n-a-1 {
				t++
				j++
			} else {
				k++
			}
		case 2:
			if j == m-a-1 {
				t++
				k--
			} else {
				j++
			}
		case 3:
			if k == a {
				t = 0
				j--
			} else {
				k--
			}
		case 0:
			if j == a+1 {
				t++
				a++
				k++
			} else {
				j--
			}
		}
	}
	return res
}
