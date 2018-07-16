package main

import "fmt"

func main() {
	fmt.Println(search([]int{3, 5, 1}, 5))
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6, 7}, 0))
	fmt.Println(search([]int{1, 2, 4, 5, 6, 7, 0}, 0))
	fmt.Println(search([]int{2, 4, 5, 6, 7, 0, 1}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 0))
	fmt.Println(search([]int{6, 7, 0, 1, 2, 4, 5}, 0))
	fmt.Println(search([]int{7, 0, 1, 2, 4, 5, 6}, 0))
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6}, 0))
	fmt.Println(search([]int{1, 2, 4, 5, 6, 0}, 0))
	fmt.Println(search([]int{2, 4, 5, 6, 0, 1}, 0))
	fmt.Println(search([]int{4, 5, 6, 0, 1, 2}, 0))
	fmt.Println(search([]int{5, 6, 0, 1, 2, 4}, 0))
	fmt.Println(search([]int{6, 0, 1, 2, 4, 5}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 8))
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6, 7}, 3))
	fmt.Println(search([]int{2}, 2))
	fmt.Println(search([]int{2}, 3))
	fmt.Println(search([]int{}, 3))
}

func search(nums []int, target int) int {
	l := len(nums)

	if l == 0 {
		return -1
	} else if l == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	s := 0
	e := l - 1
	i := (s + e) / 2

	for s < e {
		if nums[i] > nums[i+1] || i == s || i == e {
			break
		} else if nums[0] < nums[i] {
			s = i + 1
		} else {
			e = i
		}
		i = (s + e) / 2
	}

	v := bsearch(nums, 0, i, target)

	if v < 0 {
		v = bsearch(nums, i+1, l-1, target)
	}

	return v
}

func bsearch(nums []int, s int, e int, t int) int {
	if s >= e {
		if s < len(nums) && nums[s] == t {
			return s
		} else {
			return -1
		}
	}
	i := (s + e) / 2
	if t == nums[i] {
		return i
	} else if t > nums[i] {
		return bsearch(nums, i+1, e, t)
	} else {
		return bsearch(nums, s, i, t)
	}
}
