package main

import (
	"fmt"
)

func main() {
	board := make([][]byte, 4)
	board[0] = []byte{'X', 'X', 'X', 'X'}
	board[1] = []byte{'X', 'O', 'X', 'X'}
	board[2] = []byte{'O', 'X', 'O', 'X'}
	board[3] = []byte{'X', 'O', 'X', 'X'}
	solve(board)
	print(board)
}

func print(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Println()
	}
}

func solve(board [][]byte) {
	h := len(board)
	if h < 3 {
		return
	}
	w := len(board[0])
	if w < 3 {
		return
	}

	var check func(i, j int)
	check = func(i, j int) {
		if i < 0 || i == h || j < 0 || j == w {
			return
		}

		if board[i][j] == 'O' {
			board[i][j] = '1'
			check(i-1, j)
			check(i+1, j)
			check(i, j-1)
			check(i, j+1)
		}
	}

	for i := 0; i < h; i++ {
		check(i, 0)
		check(i, w-1)
	}

	for j := 0; j < w; j++ {
		check(0, j)
		check(h-1, j)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}

}
