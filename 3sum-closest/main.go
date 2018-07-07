package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(threeSumClosest([]int{-1, 2, 2, -4}, 1))
	// fmt.Println(threeSumClosest([]int{-1, 2, -4}, 1))
	// fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	// fmt.Println(threeSumClosest([]int{0, 0, 1, 2, 6, 11, 21, 333, 999990}, 19))
	// fmt.Println(threeSumClosest([]int{0, 2147483647, 2147483647, 2147483647, 2147483647}, 2147483646))
	// fmt.Println(threeSumClosest([]int{2147483648, 2147483648, 2147483648}, 2147483646))
	fmt.Println(threeSumClosest([]int{2147483648, 2147483648, 2147483648, -1, 0, 2, -2}, 2147483646))

}

func threeSumClosest(nums []int, target int) int {
	var ret int
	min := 9999999999
	l := len(nums)

	sort.Ints(nums)
	if l == 3 {
		return int(int32(nums[0] + nums[1] + nums[2]))
	}
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j, k := i+1, l-1; j < k; {
			var s int = int(int32(nums[i] + nums[j] + nums[k]))
			if s == target {
				return target
			} else if s < target {
				if min > (target - s) {
					ret = s
					min = target - s
				}
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			} else {
				if min > (s - target) {
					ret = s
					min = s - target
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return ret
}
