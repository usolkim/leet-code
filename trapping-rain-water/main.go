package main

import (
	"fmt"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(trap([]int{0, 1, 2, 2, 2, 2, 2, 2, 1, 0}))
	fmt.Println(trap([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(trap([]int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}))
	fmt.Println(trap([]int{6, 5, 4, 3, 3, 2, 1, 1, 0}))
	fmt.Println(trap([]int{6, 5, 4, 3, 2, 1, 6, 0}))
	fmt.Println(trap([]int{11, 1, 20, 0, 33, 2, 54, 11}))
	fmt.Println(trap([]int{55, 1, 20, 0, 33, 2, 54, 11}))
	fmt.Println(trap([]int{0}))
	fmt.Println(trap([]int{2, 10, 2}))
}

func trap(height []int) int {
	l := len(height)
	if l < 3 {
		return 0
	}
	left := 0
	right := l - 1
	max := 0

	for ; left < l-1 && height[left] <= height[left+1]; left++ {
	}

	for ; right > 0 && height[right-1] >= height[right]; right-- {
	}

	res := 0
	t := 'l'

	if height[left] < height[right] {
		t = 'l'
		max = height[left]
	} else {
		t = 'r'
		max = height[right]
	}

	for left < right {
		if height[left] < height[right] {
			if t == 'l' {
				if max < height[left] {
					max = height[left]
				} else {
					res += max - height[left]
				}
			} else {
				t = 'l'
				max = height[left]
			}
			left++
		} else {
			if t == 'r' {
				if max < height[right] {
					max = height[right]
				} else {
					res += max - height[right]
				}
			} else {
				t = 'r'
				max = height[right]
			}
			right--
		}
	}

	return res
}
