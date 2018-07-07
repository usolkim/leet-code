package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("1"))
	fmt.Println(letterCombinations("212"))
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("34"))
	fmt.Println(letterCombinations("349"))
}

func letterCombinations(digits string) []string {
	l := len(digits)
	switch l {
	case 0:
		return nil
	case 1:
		return getLetters(digits[0])
	case 2:
		lhs := getLetters(digits[0])
		rhs := getLetters(digits[1])
		return combine(lhs, rhs)
	default:
		lhs := letterCombinations(digits[0 : l/2])
		rhs := letterCombinations(digits[l/2 : l])
		return combine(lhs, rhs)

	}
}

func combine(lhs []string, rhs []string) []string {
	var ret []string
	for _, lv := range lhs {
		for _, rv := range rhs {
			ret = append(ret, lv+rv)
		}
	}
	return ret
}

func getLetters(digit byte) []string {
	switch digit {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	}
	return nil
}
