package main

import "fmt"

func main() {
	grid := make([][]int, 3)
	grid[0] = []int{1, 3, 1}
	grid[1] = []int{1, 5, 1}
	grid[2] = []int{4, 2, 1}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	arr := make([]int, cols)
	
	arr[0] = grid[0][0]
	for i:=1;i<cols;i++{
		arr[i] = arr[i-1] + grid[0][i]
	}

	for i:=1;i<rows;i++{
		arr[0] +=grid[i][0]
		for j:=1;j<cols;j++{
			if arr[j-1] < arr[j] {
				arr[j] = arr[j-1] + grid[i][j]
			}else{
				arr[j] = arr[j] + grid[i][j]
			}
		}
	}

	return arr[cols-1]
}
