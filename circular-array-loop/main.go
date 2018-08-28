package main

import "fmt"

func main(){
	fmt.Println(circularArrayLoop([]int{-2, -1, 1, -2, 2}))
	fmt.Println(circularArrayLoop([]int{-1,2}))
	fmt.Println(circularArrayLoop([]int{4,1,1,1}))
	fmt.Println(circularArrayLoop([]int{4,1,1,2}))
	fmt.Println(circularArrayLoop([]int{5,1,1,1}))
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
	fmt.Println(circularArrayLoop([]int{2, -1, 1, -2, 3}))
}

func circularArrayLoop(nums []int) bool {
    l := len(nums)
    if l < 2 {
        return false
    }
    sum:=0
    i:=0
    for j:=0;j<l;j++ {
        if nums[i] == l {
            i = (i + 1) % l
            continue
        }
        sum += nums[i]
        //fmt.Println(j, i, sum, nums)
        if sum != 0 && sum % l == 0 {
            return true
        }
        next := (i+nums[i])% l
        if next < 0 {
            next += l
        }
        i = next
        //fmt.Println("next:", next)
    }
    return false
}