package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isNumber(" "))
	fmt.Println(isNumber("-"))
	fmt.Println(isNumber("+"))
	fmt.Println(isNumber("-+1"))
	fmt.Println(isNumber("-e134"))
	fmt.Println(isNumber("e5"))
	fmt.Println(isNumber("e+5"))
	fmt.Println(isNumber("-a"))
	fmt.Println(isNumber("-1.4s3"))
	fmt.Println(isNumber("0. 41"))
	fmt.Println(isNumber("."))
	fmt.Println(isNumber(".e1"))
	fmt.Println(isNumber("-."))
	fmt.Println(isNumber("6e6.5"))
	fmt.Println("==============================")
	fmt.Println(isNumber("1"))
	fmt.Println(isNumber("-1"))
	fmt.Println(isNumber("+1"))
	fmt.Println(isNumber("-0"))
	fmt.Println(isNumber("+0"))
	fmt.Println(isNumber("0.123"))
	fmt.Println(isNumber(".123"))
	fmt.Println(isNumber("-.123"))
	fmt.Println(isNumber("-10."))
	fmt.Println(isNumber("+1.4725"))
	fmt.Println(isNumber("1e0"))
	fmt.Println(isNumber("1.4e5"))
	fmt.Println(isNumber("+1e0"))
	fmt.Println(isNumber("-1e0"))
	fmt.Println(isNumber("-1e-5"))
	fmt.Println(isNumber("-1e+5"))
	fmt.Println(isNumber("1.0e5"))
	fmt.Println(isNumber("1.e5"))
	fmt.Println(isNumber("1123411e-15"))
	fmt.Println(isNumber("-1123411e+15"))
	fmt.Println(isNumber(" -1123411e+15"))
	fmt.Println(isNumber("-1123411e+15 "))
	fmt.Println(isNumber(" -1123411e+15 "))
}

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if len(s) == 0 || s == "." {
		return false
	}
	dot := false
	e := false
	for i, c := range s {
		var a, b byte
		if i > 0 {
			b = s[i-1]
		}
		if i < len(s)-1 {
			a = s[i+1]
		}
		if c == '+' || c == '-' {
			if i > 0 && b != 'e' || i == len(s)-1 {
				return false
			}
		} else if c == '.' {
			if e || dot || ((b < '0' || b > '9') && (a < '0' || a > '9')) {
				return false
			}
			dot = true
		} else if c == 'e' {
			if e || i == 0 || i > 0 && b != '.' && (b < '0' || b > '9') || i == 1 && b == '.' || i == len(s)-1 {
				return false
			}
			e = true
		} else if c >= '0' && c <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
