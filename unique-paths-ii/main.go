package main

import (
	"fmt"
)

func main() {
	obstacleGrid := make([][]int, 3)
	obstacleGrid[0] = []int{0, 0, 0}
	obstacleGrid[1] = []int{1, 0, 1}
	obstacleGrid[2] = []int{0, 0, 0}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	if row == 0 {
		return 0
	}

	col := len(obstacleGrid[0])
	if col == 0 {
		return 0
	}

	arr := make([]int, col)
	arr[0] = 1
	for i := 0; i < row; i++ {
		if obstacleGrid[i][0] != 0 {
			arr[0] = 0
		}
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 0 {
				arr[j] = arr[j] + arr[j-1]
			} else {
				arr[j] = 0
			}
		}
	}

	return arr[col-1]
}
