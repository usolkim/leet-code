package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{"", "b"}))
	fmt.Println(longestCommonPrefix([]string{"abcde"}))
	fmt.Println(longestCommonPrefix([]string{"abcdeasdfasdfasdf", "abcdeadfasdfasdf", "ab"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car", "dog", "racecar", "car", "dog", "racecar", "car", "dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return strs[0]
	}
	i := 0
	for i < len(strs[0]) {
		for j := 1; j < l; j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				goto OUT
			}
		}
		i++
	}

OUT:
	return strs[0][0:i]
}
