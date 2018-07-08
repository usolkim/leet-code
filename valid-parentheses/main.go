package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid(""))
	fmt.Println(isValid("{}"))
	fmt.Println(isValid("[]"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("(("))
	fmt.Println(isValid("))"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	a := make([]byte, 0)
	for i := 0; i < l; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			a = append(a, s[i])
		} else {
			if len(a) == 0 {
				return false
			}
			b := a[len(a)-1]
			if b == '(' && s[i] == ')' ||
				b == '{' && s[i] == '}' ||
				b == '[' && s[i] == ']' {
				a = a[0 : len(a)-1]
			} else {
				return false
			}
		}
	}
	if len(a) == 0 {
		return true
	}
	return false
}
