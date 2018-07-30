package main

import (
	"fmt"
)

func main() {
	board := make([][]byte, 3)
	board[0] = []byte{'A', 'B', 'C', 'E'}
	board[1] = []byte{'S', 'F', 'C', 'S'}
	board[2] = []byte{'A', 'D', 'E', 'E'}
	fmt.Println(exist(board, ""))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "SEECCE"))
	fmt.Println(exist(board, "SFDEES"))
	fmt.Println(exist(board, "ABCESEEDASFC"))

	fmt.Println(exist(board, "SEECS"))
	fmt.Println(exist(board, "SFDAS"))

}

func exist(board [][]byte, word string) bool {
	l := len(word)
	if l == 0 {
		return true
	}
	row := len(board)
	if row == 0 {
		return false
	}
	col := len(board[0])
	if col == 0 {
		return false
	}

	check := make([][]bool, row)
	for i := 0; i < row; i++ {
		check[i] = make([]bool, col)
	}

	var backt func(int, int, int) bool
	backt = func(i, j, w int) bool {
		if check[i][j] {
			return false
		}
		if board[i][j] != word[w] {
			return false
		}
		if l-1 == w {
			return true
		}
		check[i][j] = true
		if i > 0 && backt(i-1, j, w+1) {
			return true
		} else if i < row-1 && backt(i+1, j, w+1) {
			return true
		} else if j > 0 && backt(i, j-1, w+1) {
			return true
		} else if j < col-1 && backt(i, j+1, w+1) {
			return true
		}
		check[i][j] = false
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if backt(i, j, 0) {
				return true
			}
		}
	}
	return false
}
