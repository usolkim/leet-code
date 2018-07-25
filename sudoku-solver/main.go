package main

import (
	"fmt"
	"math"
)

func main() {
	board := make([][]byte, 9)
	board[0] = []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
	board[1] = []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	board[2] = []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	board[3] = []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	board[4] = []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	board[5] = []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	board[6] = []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	board[7] = []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	board[8] = []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}

	solveSudoku(board)

	print(board)
}

func print(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				fmt.Print("_ ")
			} else {
				fmt.Print(board[i][j]-'0', " ")
			}
		}
		fmt.Println()
	}
}

func solveSudoku(board [][]byte) {
	row := make([]int, 9)
	col := make([]int, 9)
	areas := make([]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				area := i/3*3 + j/3
				n := board[i][j] - '0'
				nb := 1 << (n - 1)
				row[i] = row[i] | nb
				col[j] = col[j] | nb
				areas[area] = areas[area] | nb
			}
		}
	}

	backt(0, board, row, col, areas, 1<<9-1)
}

func backt(index int, board [][]byte, row, col, areas []int, f int) bool{
	if index == 81 {
		return true
	}
	i := index / 9
	j := index % 9
	if board[i][j] == '.' {
		area := i/3*3 + j/3

		avail := f & ^(row[i] | col[j] | areas[area])

		for avail > 0 {
			pos := avail & (^avail + 1)
			avail -= pos
			n := math.Log2(float64(pos))
			board[i][j] = byte('1' + int(n))
			row[i] = row[i] | pos
			col[j] = col[j] | pos
			areas[area] = areas[area] | pos
			res := backt(index+1, board, row, col, areas, f)
			if res {
				return true
			}
			board[i][j] = '.'
			row[i] = row[i] & ^pos
			col[j] = col[j] & ^pos
			areas[area] = areas[area] & ^pos
		}
		return false
	} 
		return backt(index+1, board, row, col, areas, f)
}
