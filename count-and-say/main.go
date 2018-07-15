package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(0))
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
	fmt.Println(countAndSay(20))
}

func countAndSay(n int) string {
	if n < 2 {
		return "1"
	}
	return recursive("1", 1, n)
}

func recursive(s string, c int, n int) string {
	if c == n {
		return s
	}
	var ss string
	st := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			ss = ss + string('0'+i-st+1) + string(s[i])
			st = i + 1
		}
	}

	ss = ss + string('0'+len(s)-st) + string(s[len(s)-1])

	return recursive(ss, c+1, n)
}
