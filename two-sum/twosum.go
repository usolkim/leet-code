package main

import (
    "fmt"
)

func main(){
    fmt.Println("main")
    var nums = []int{3, 1, 7, 11, 2, 15}
    fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
    var hash = map[int]int{}
    var result []int
    for i, num := range nums {
        if j, exist := hash[target-num]; exist {
            result = append(result, i, j)
            break
        }else{
            hash[num] = i
        }
    }
    return result
}