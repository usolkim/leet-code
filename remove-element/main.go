package main

import "fmt"

func main() {
	nums := []int{}
	l := removeElement(nums, 1)

	for i := 0; i < l; i++ {
		fmt.Print(nums[i])
	}
	fmt.Println()

}

func removeElement(nums []int, val int) int {
	c := 0
	for i := range nums {
		if nums[i] != val {
			nums[c] = nums[i]
			c++
		}
	}
	return c
}
