package main

import "fmt"

func main() {
	matrix := make([][]int, 0)
	print(matrix)
	fmt.Println(searchMatrix(matrix, 3))

	matrix = make([][]int, 1)
	matrix[0] = []int{}
	print(matrix)
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 13))

	matrix = make([][]int, 1)
	matrix[0] = []int{1}
	print(matrix)
	fmt.Println(searchMatrix(matrix, 3))

	matrix = make([][]int, 2)
	matrix[0] = []int{}
	matrix[1] = []int{}
	print(matrix)
	fmt.Println(searchMatrix(matrix, 0))

	matrix = make([][]int, 2)
	matrix[0] = []int{1}
	matrix[1] = []int{3}
	print(matrix)
	fmt.Println(searchMatrix(matrix, 1))
	fmt.Println(searchMatrix(matrix, 2))
	fmt.Println(searchMatrix(matrix, 3))

	matrix = make([][]int, 4)
	matrix[0] = []int{1, 3, 5, 7}
	matrix[1] = []int{10, 11, 16, 20}
	matrix[2] = []int{23, 30, 34, 50}
	matrix[3] = []int{123, 130, 134, 150}
	print(matrix)
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 13))

	matrix = make([][]int, 4)
	matrix[0] = []int{1, 3, 5, 7, 9}
	matrix[1] = []int{10, 11, 16, 20, 21}
	matrix[2] = []int{23, 30, 34, 50, 100}
	matrix[3] = []int{123, 130, 134, 150}
	print(matrix)
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 13))
	fmt.Println(searchMatrix(matrix, 99))

	matrix = make([][]int, 3)
	matrix[0] = []int{1, 3, 5, 7}
	matrix[1] = []int{10, 11, 16, 20}
	matrix[2] = []int{23, 30, 34, 50}
	print(matrix)
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 13))

	matrix[0] = []int{1, 3, 5, 7, 8}
	matrix[1] = []int{10, 11, 16, 20, 22}
	matrix[2] = []int{23, 30, 34, 43, 50}
	print(matrix)
	fmt.Println(searchMatrix(matrix, 0))
	fmt.Println(searchMatrix(matrix, 9))
	fmt.Println(searchMatrix(matrix, 21))
	fmt.Println(searchMatrix(matrix, 31))
	fmt.Println(searchMatrix(matrix, 44))
	fmt.Println(searchMatrix(matrix, 51))
	fmt.Println(searchMatrix(matrix, 1))
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 8))
	fmt.Println(searchMatrix(matrix, 11))
	fmt.Println(searchMatrix(matrix, 20))
	fmt.Println(searchMatrix(matrix, 22))
	fmt.Println(searchMatrix(matrix, 30))
	fmt.Println(searchMatrix(matrix, 43))

}

func print(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("-------------------")
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	return rbsearch(matrix, 0, len(matrix)-1, target)
}

func rbsearch(matrix [][]int, s, e, target int) bool {
	if s+1 >= e {
		if matrix[s][0] == target || matrix[e][0] == target {
			return true
		} else if len(matrix[e]) == 1 {
			return false
		} else if matrix[e][0] < target {
			return cbsearch(matrix[e], 1, len(matrix[e])-1, target)
		}
		return cbsearch(matrix[s], 1, len(matrix[s])-1, target)
	}
	c := (s + e) / 2
	if matrix[c][0] == target {
		return true
	} else if matrix[c][0] < target {
		return rbsearch(matrix, c, e, target)
	}
	return rbsearch(matrix, s, c-1, target)
}

func cbsearch(row []int, s, e, target int) bool {
	if s == e {
		return row[s] == target
	}
	c := (s + e) / 2
	if row[c] == target {
		return true
	} else if row[c] < target {
		return cbsearch(row, c+1, e, target)
	}
	return cbsearch(row, s, c, target)
}
