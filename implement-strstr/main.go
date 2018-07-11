package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("hell", "ll"))
	fmt.Println(strStr("aaaa", "bb"))
	fmt.Println(strStr("aaaa", ""))
	fmt.Println(strStr("aaaa", "a"))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("", "a"))
	fmt.Println(strStr("abcbcbc", "bc"))
	fmt.Println(strStr("abc", "abc"))
}

func strStr(haystack string, needle string) int {
	hl := len(haystack)
	nl := len(needle)
	if nl == 0 {
		return 0
	}
	if hl == 0 || hl < nl {
		return -1
	}
	for i := 0; i <= hl-nl; i++ {
		if string(haystack[i:i+nl]) == needle {
			return i
		}
	}
	return -1
}
