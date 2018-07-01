package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("c"))
	fmt.Println(lengthOfLongestSubstring("au"))
}

func lengthOfLongestSubstring(s string) int {
	l := len(s)

	if l < 2 {
		return l
	}

	table := [128]int{}

	for i := 0; i < 128; i++ {
		table[i] = -1
	}

	var result, start int = 0, 0

	for end, rune := range s {
		next := int(rune)
		index := table[next]
		if index >= start {
			start = index + 1
		}
		table[next] = end
		if result < end-start+1 {
			result = end - start + 1
		}
	}

	return result
}
