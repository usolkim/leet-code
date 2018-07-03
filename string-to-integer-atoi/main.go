package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(myAtoi("-000000000000001"))
	fmt.Println(myAtoi("  0000000000012345678"))

	fmt.Println(myAtoi("0-1"))
	fmt.Println(myAtoi("20000000000000000000"))
	fmt.Println(myAtoi("-+1"))
	fmt.Println(myAtoi("+-1"))
	fmt.Println(myAtoi("+"))
	fmt.Println(myAtoi("++1"))
	fmt.Println(myAtoi("+ 1"))
	fmt.Println(myAtoi(" +1"))
	fmt.Println(myAtoi("+1l"))
	fmt.Println(myAtoi("+d1"))
	fmt.Println(myAtoi("+1"))
	fmt.Println(myAtoi("-"))
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("1"))
	fmt.Println(myAtoi("-1"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("   - 42"))
	fmt.Println(myAtoi("   -a42"))
	fmt.Println(myAtoi("     1l1361065549"))
	fmt.Println(myAtoi("09"))
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("91283472332"))
}

func myAtoi(str string) int {
	var dirty bool
	var start, end, symbol int = -1, -1, 0
	// var start, end, minus, plus int = -1, -1, 1, 1
	for i, ch := range str {
		if !dirty {
			switch {
			case symbol == 0 && ch == 32:
				continue
			case symbol == 0 && (ch == 43 || ch == 45):
				symbol = 44 - int(ch)
			case symbol == 0 && ch == 48:
				dirty = true
			case ch > 48 && ch <= 57:
				start = i
				end = i
				dirty = true
			default:
				return 0
			}
		} else if ch < 48 || ch > 57 {
			break
		} else {
			end++
		}
	}

	if start < 0 {
		return 0
	}

	var max_32, min_32 int = 1<<31 - 1, -1 << 31

	if end-start > 10 {
		if symbol == -1 {
			return min_32
		} else {
			return max_32
		}
	}

	v := str[start : end+1]

	res, err := strconv.ParseInt(v, 10, 64)

	if err != nil {
		return 0
	}

	res *= int64(symbol)

	if res < int64(min_32) {
		return min_32
	} else if res > int64(max_32) {
		return max_32
	}

	return int(res)
}
