package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aaaaaaabaaababbbabbbaabbaabbaaaaaabaabbaaababababab", "ba******"))
	fmt.Println(isMatch("mississippi", "m??*ss*?i*pi"))
	fmt.Println(isMatch("ho", "**ho"))
	fmt.Println(isMatch("", "a"))
	fmt.Println(isMatch("", ""))
	fmt.Println(isMatch("a", ""))
	fmt.Println(isMatch("a", "?"))
	fmt.Println(isMatch("a", "*"))
	fmt.Println(isMatch("", "*"))
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("cb", "?a"))
	fmt.Println(isMatch("adceb", "*a*b"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
}

func isMatch(s string, p string) bool {
	ls := len(s)
	lp := len(p)

	if ls == 0 && lp == 0 {
		return true
	}
	if lp == 0 {
		return false
	}

	i := 0
	j := 0
	stari := -1
	starj := 0

	for i < ls {
		if j < lp && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < lp && p[j] == '*' {
			stari = i
			j++
			starj = j
		} else if stari != -1 {
			stari++
			i = stari
			j = starj
		} else {
			return false
		}
	}

	for ; j < lp; j++ {
		if p[j] != '*' {
			return false
		}
	}

	return true
}
