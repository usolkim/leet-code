package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{1}, 2))
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2, 3, 5}, 1))
	fmt.Println(combinationSum([]int{2, 4, 6}, 13))
	fmt.Println(combinationSum([]int{3}, 4))
	fmt.Println(combinationSum([]int{}, 4))
}

func combinationSum(candidates []int, target int) [][]int {
	if target <= 0 {
		return nil
	}
	l := len(candidates)
	if l == 0 {
		return nil
	}
	if l == 1 {
		if target%candidates[0] != 0 {
			return nil
		}
	}

	ret := make([][]int, 0)

	for i, v := range candidates {
		if v == target {
			ret = append(ret, []int{v})
		}
		r := sum(candidates[i:l], target-v, v)
		ret = append(ret, r...)
	}

	return ret
}

func sum(c []int, target int, p int) [][]int {
	r := make([][]int, 0)
	for i, v := range c {
		if v == target {
			r = append(r, []int{p, v})
		} else if v < target {
			sr := sum(c[i:len(c)], target-v, v)
			for _, rv := range sr {
				tmp := []int{p}
				tmp = append(tmp, rv...)
				r = append(r, tmp)
			}
		}
	}
	return r
}
