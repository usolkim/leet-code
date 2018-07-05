package main

import "fmt"

func main() {
	fmt.Println(isMatch("", ".*"))
	fmt.Println(isMatch("", ""))
	fmt.Println(isMatch("a", ""))
	fmt.Println(isMatch("", "c*"))
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("aabdc", ".*c"))
	fmt.Println(isMatch("cdbd", ".*c.*"))
	fmt.Println(isMatch("aaacdbd", ".*c.*"))
	fmt.Println(isMatch("aaacccdbd", ".*c.*"))
	fmt.Println(isMatch("cccccc", ".c*c"))
}

func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)
	if lenP == 0 {
		return lenS == 0
	}

	memo := make([][]bool, lenS+1)
	for i := 0; i <= lenS; i++ {
		memo[i] = make([]bool, lenP+1)
	}
	memo[lenS][lenP] = true

	for i := lenS; i >= 0; i-- {
		for j := lenP - 1; j >= 0; j-- {
			f := i < lenS && (p[j] == '.' || s[i] == p[j])
			if j < lenP-1 && p[j+1] == '*' {
				memo[i][j] = memo[i][j+2] || (f && memo[i+1][j])
			} else {
				memo[i][j] = f && memo[i+1][j+1]
			}
		}
	}
	return memo[0][0]
}
