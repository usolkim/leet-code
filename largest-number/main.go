package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}))
	fmt.Println(largestNumber([]int{121, 12}))
	fmt.Println(largestNumber([]int{12, 121}))
}

type StringSlice []string

func (strs StringSlice) Len() int { return len(strs) }
func (strs StringSlice) Less(i, j int) bool {
	a, _ := strconv.Atoi(strs[i] + strs[j])
	b, _ := strconv.Atoi(strs[j] + strs[i])
	return a > b
}
func (strs StringSlice) Swap(i, j int) { strs[i], strs[j] = strs[j], strs[i] }

func largestNumber(nums []int) string {
	l := len(nums)
	if l == 0 {
		return "0"
	}

	strs := make([]string, l)
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	sort.Sort(StringSlice(strs))

	if strs[0] == "0" && strs[l-1] == "0" {
		return "0"
	}

	ret := ""
	for _, s := range strs {
		ret = ret + s
	}
	return ret
}
