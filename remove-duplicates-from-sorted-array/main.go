package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5, 5, 5, 6}
	l := removeDuplicates(nums)

	for i := 0; i < l; i++ {
		fmt.Print(nums[i])
	}
	fmt.Println()

}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l > 1 {
		for i := 0; i < l-1; {
			if nums[i] == nums[i+1] {
				if i+2 < l {
					nums = append(nums[0:i+1], nums[i+2:l]...)
				} else {
					nums = nums[0 : i+1]
				}
				l--
			} else {
				i++
			}
		}
	}
	return len(nums)
}
