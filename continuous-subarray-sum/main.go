package main

import "fmt"

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	fmt.Println(checkSubarraySum([]int{24, 2, 5, 4, 6}, 6))
	fmt.Println(checkSubarraySum([]int{24, 2, 5, 4, 6}, 0))
	fmt.Println(checkSubarraySum([]int{0, 0}, 0))
}

func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}

	if k == 0 {
		for _, n := range nums {
			if n > 0 {
				return false
			}
		}
		return true
	}

	memo := make(map[int]int, 0)
	memo[0] = -1

	sum := 0

	for i, n := range nums {
		sum += n
		remain := sum % k
		if j, exist := memo[remain]; exist {
			if i-j >= 2 {
				return true
			}
		} else {
			memo[remain] = i
		}
	}

	return false
}
