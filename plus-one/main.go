package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func plusOne(digits []int) []int {
	l := len(digits)
	digits[l-1]++
	for i := l - 1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i-1]++
			digits[i] = 0
			continue
		}
		return digits
	}
	if digits[0] == 10 {
		digits[0] = 0
		res := make([]int, 0)
		res = append(res, 1)
		res = append(res, digits...)
		return res
	}
	return digits
}
