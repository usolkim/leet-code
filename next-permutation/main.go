package main

import "fmt"

func main() {
	var nums []int

	nums = []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 5, 8, 3, 8, 7, 6, 5, 4}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 5, 8, 3, 7, 6, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 1, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	i := len(nums) - 1
	for i > 0 {
		i--
		if nums[i] < nums[i+1] {
			j := i + 1
			for ; j < len(nums); j++ {
				if nums[i] >= nums[j] {
					break
				}
			}

			swap(nums, i, j-1)

			i++
			break
		}
	}

	for k := 0; k < (len(nums)-i)/2; k++ {
		swap(nums, i+k, len(nums)-k-1)
	}
}

func swap(nums []int, i int, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
