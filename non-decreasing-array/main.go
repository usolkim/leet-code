package main

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{1, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{2, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{1, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{2, 4, 1, 3}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 5}))
	fmt.Println(checkPossibility([]int{1}))
	fmt.Println(checkPossibility([]int{}))
	fmt.Println(checkPossibility([]int{1,1,1,1,1}))
	fmt.Println(checkPossibility([]int{1,2,3,4,5}))
	fmt.Println(checkPossibility([]int{1,1,2,3,3,3,4,4,5,5,5,5}))
	fmt.Println(checkPossibility([]int{1,1,2,3,3,3,4,3,5,5,5,5}))
	fmt.Println(checkPossibility([]int{5,4,3,2,1}))
}

func checkPossibility(nums []int) bool {
	check := false
	for i:=0;i<len(nums)-1;i++{
		if nums[i] > nums[i+1] {
			if check {
				return false
			}
			check = true
			if i == 0 || nums[i-1] <= nums[i+1]{
				nums[i] = nums[i+1]
			}else{
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}
