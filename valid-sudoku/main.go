package main

import "fmt"

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

	fmt.Println(isValidSudoku(board))

	board[0][0] = '8'
	fmt.Println(isValidSudoku(board))

}

func isValidSudoku(board [][]byte) bool {
	rows := make([][]int, 9)
	cols := make([][]int, 9)
	areas := make([][]int, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make([]int, 9)
		cols[i] = make([]int, 9)
		areas[i] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			n := int(board[i][j]-'0') - 1
			if rows[i][n] > 0 {
				return false
			}
			rows[i][n]++

			if cols[n][j] > 0 {
				return false
			}
			cols[n][j]++

			area := i/3*3 + j/3
			if areas[area][n] > 0 {
				return false
			}
			areas[area][n]++
		}
	}
	return true
}
