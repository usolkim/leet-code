package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(0))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	if n < 0 {
		return nil
	}
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	return recursive(n, n)
}

func recursive(o int, c int) []string {
	ret := make([]string, 0)
	for i := o; i > 0; i-- {
		for j := 1; o == c && j <= i || o < c && j <= c; j++ {
			if o-i == 0 && c-j == 0 {
				ret = append(ret, getString(i, j))
			}
			if o-i > 0 && c-j > 0 {
				r := recursive(o-i, c-j)
				s := getString(i, j)
				for _, v := range r {
					ret = append(ret, s+v)
				}
			}
		}
	}
	return ret
}

func getString(i int, j int) string {
	s := make([]rune, 0)
	for k := 0; k < i; k++ {
		s = append(s, '(')
	}
	for k := 0; k < j; k++ {
		s = append(s, ')')
	}
	return string(s)
}

// 1
// () 1 -1

// 2
// (()) 2 -2
// ()() 1 -1 1 -1

// 3
// ((())) 3 -3
// (()()) 2 -1 1 -2
// (())() 2 -2 1 -1
// ()(()) 1 -1 2 -2
// ()()() 1 -1 1 -1 1 -1

// 4
// (((()))) 4 -4
// ((()())) 3 -1 1 -3
// ((())()) 3 -2 1 -2
// ((()))() 3 -3 1 -1
// (()(())) 2 -1 2 -3
// (()()()) 2 -1 1 -1 1 -2
// (()())() 2 -1 1 -2 1 -1
// (())(()) 2 -2 2 -2
// (())()() 2 -2 1 -1 1 -1
// ()((())) 1 -1 3 -3
// ()(()()) 1 -1 2 -1 1 -2
// ()(())() 1 -1 2 -2 1 -1
// ()()(()) 1 -1 1 -1 2 -2
// ()()()() 1 -1 1 -1 1 -1 1 -1
