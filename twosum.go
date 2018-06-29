package main

import (
    "fmt"
)

func main(){
    var nums = []int{3, 1, 7, 11, 2, 15}
    fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
    var hash = map[int]int{}
    for i, num := range nums {
        if j, exist := hash[target-num]; exist {
            return []int{j,i}
        }else{
            hash[num] = i
        }
    }
    return nil
}