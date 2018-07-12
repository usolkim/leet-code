package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses("("))
	fmt.Println(longestValidParentheses(")"))
	fmt.Println(longestValidParentheses("(("))
	fmt.Println(longestValidParentheses("()"))
	fmt.Println(longestValidParentheses("))"))
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(")()())()()("))
	fmt.Println(longestValidParentheses("(()))(()(()())"))
	fmt.Println(longestValidParentheses("()(()())))((()(((()))()))()()((()()))"))
	fmt.Println(longestValidParentheses("()(()())))((()(((()))()))()()((()())("))
}

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	ret := 0

	c := make([]int, len(s))

	d := 0
	for i, ss := range s {
		if ss == '(' {
			if i > 0 && c[i-1] < 0 {
				d = 1
			} else {
				d++
			}
		} else {
			d--
		}
		c[i] = d
	}

	start := -1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			if start < 0 {
				continue
			} else {
				if c[i] == c[start] {
					start = -1
				} else if c[i] == c[start]+1 && ret < start-i+1 {
					ret = start - i + 1
				}
			}
		} else {
			if start < 0 && c[i] >= 0 {
				start = i
			} else if c[i] < 0 {
				start = -1
			}
		}
	}

	return ret
}
