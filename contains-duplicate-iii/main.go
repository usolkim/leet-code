package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 0, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{}, 0, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1}, 0, 1))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1}, 1, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 111, 222, 3, 333}, 2, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 111, 222, 3, 333}, 3, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 111, 222, 3, 333}, 3, 1))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 111, 222, 3, 333}, 3, 3))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := len(nums)
	if l == 0 || k < 1 || t < 0 {
		return false
	}

	if k > l {
		k = l
	}

	for i := 1; i <= k; i++ {
		for j := 0; j+i < l; j++ {
			diff := int(math.Abs(float64(nums[j] - nums[j+i])))
			if t >= diff {
				return true
			}
		}
	}
	return false
}
