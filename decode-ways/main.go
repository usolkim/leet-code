package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings(""))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("00"))
	fmt.Println(numDecodings("000"))
	fmt.Println(numDecodings("100"))
	fmt.Println(numDecodings("110"))
	fmt.Println(numDecodings("1"))
	fmt.Println(numDecodings("10"))
	fmt.Println(numDecodings("121212"))
	fmt.Println(numDecodings("10192631"))
	fmt.Println(numDecodings("333333"))
}

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	ar := make([]int, l+1)

	ar[l] = 1

	if s[l-1] != '0' {
		ar[l-1] = 1
	}

	for i := l - 2; i >= 0; i-- {
		n := s[i] - '0'
		if n == 0 {
			if ar[i+1] == 0 {
				return 0
			}
		} else if n == 1 || n == 2 && s[i+1]-'0' < 7 {
			ar[i] = ar[i+1] + ar[i+2]
		} else {
			ar[i] = ar[i+1]
		}
	}
	return ar[0]
}
