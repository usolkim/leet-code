package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 2, 3}))
	fmt.Println(maxArea([]int{3, 2, 1}))
	fmt.Println(maxArea([]int{3, 1, 2}))
	fmt.Println(maxArea([]int{30, 29, 4, 627, 93}))
}

func maxArea(height []int) int {
	var i, j, max int = 0, len(height) - 1, 0

	for i < j {
		var area int
		if height[i] < height[j] {
			area = (j - i) * height[i]
			i++
		} else {
			area = (j - i) * height[j]
			j--
		}
		if max < area {
			max = area
		}
	}
	return max
}
