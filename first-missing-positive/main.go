package main

import (
	"fmt"
)

func main(){
	fmt.Println(firstMissingPositive([]int{1,2,0}))
	fmt.Println(firstMissingPositive([]int{3,4,-1,1}))
	fmt.Println(firstMissingPositive([]int{7,8,9,11,12}))
	fmt.Println(firstMissingPositive([]int{5,1,4,3,2,8,7}))
	fmt.Println(firstMissingPositive([]int{1,2,3,4,5,6,7,8}))
	fmt.Println(firstMissingPositive([]int{6,5,4,3,2,1}))
}

func firstMissingPositive(nums []int) int {
	l := len(nums)
	a := make([]int, l)
	for i:=0;i<l;i++{
		n := nums[i]
		if n < 1 || n> l {
			continue
		}
		a[n-1] = 1
	}
	for i:=0;i<l;i++{
		if a[i] == 0 {
			return i+1
		}
	}
    return l+1
}