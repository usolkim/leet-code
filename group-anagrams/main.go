package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	l := len(strs)
	if l == 0 {
		return nil
	} else if l == 1 {
		return [][]string{strs}
	}

	m := make(map[string][]string, 0)

	for _, s := range strs {
		n := make([]int, len(s))
		for i, r := range s {
			n[i] = int(r)
		}
		sort.Ints(n)
		runes := make([]rune, len(s))
		for i, r := range n {
			runes[i] = rune(r)
		}
		sr := string(runes)
		if v, e := m[sr]; e {
			m[sr] = append(v, s)
		} else {
			m[sr] = []string{s}
		}
	}

	res := make([][]string, 0)

	for _, s := range m {
		res = append(res, s)
	}
	return res
}
