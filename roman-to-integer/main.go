package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	l := len(s)

	ret := 0

	for i := l - 1; i >= 0; i-- {
		switch s[i] {
		case 'I':
			if ret >= 5 {
				ret--
			} else {
				ret++
			}
		case 'V':
			ret += 5
		case 'X':
			if ret >= 50 {
				ret -= 10
			} else {
				ret += 10
			}
		case 'L':
			ret += 50
		case 'C':
			if ret >= 500 {
				ret -= 100
			} else {
				ret += 100
			}
		case 'D':
			ret += 500
		case 'M':
			ret += 1000
		}
	}
	return ret
}
