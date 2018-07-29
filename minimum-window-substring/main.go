package main

import (
	"fmt"
)

func main() {
	fmt.Println("\"" + minWindow("ADOBECODEBANC", "ABC") + "\"")
	fmt.Println("\"" + minWindow("ADOBEBODCAN", "ABC") + "\"")
	fmt.Println("\"" + minWindow("ADOBEBODCAN", "ABCD") + "\"")
	fmt.Println("\"" + minWindow("ADOBEBODCAN", "") + "\"")
	fmt.Println("\"" + minWindow("", "ABC") + "\"")
	fmt.Println("\"" + minWindow("ADOBECODEBANC", "ABCZ") + "\"")
	fmt.Println("\"" + minWindow("bbaa", "aba") + "\"")
}

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	if len(s) < len(t) {
		return ""
	}
	if s == t {
		return s
	}

	tmap := [128]int{}
	for _, v := range t {
		tmap[byte(v)]++
	}
	start := 0
	curr := 0
	next := 1
	count := len(t)
	res := ""
	for start < len(s) {
		if curr < next {
			if tmap[s[curr]] > 0 {
				count--
			}
			tmap[s[curr]]--
		}
		if count == 0 {
			if res == "" || len(res) > curr-start+1 {
				res = s[start : curr+1]
			}
			if tmap[s[start]] == 0 {
				count++
			}
			tmap[s[start]]++
			start++
			next = curr
		} else if curr < len(s)-1 {
			curr++
			next = curr + 1
		} else {
			break
		}

	}
	return res
}
