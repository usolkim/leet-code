package main

import (
	"fmt"
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
	var allow_blank, zeros bool = true, true
	var start, end, symbol int = -1, -1, 0

	for i, ch := range str {
		switch {
			case ch == 32 :
				if !allow_blank {
					goto OUT
				}
			case ch == 43 || ch == 45 :
				if symbol == 0 {
					allow_blank = false
					symbol = 44 - int(ch)
				} else {
					goto OUT
				}
			case ch == 48 :
				if symbol == 0 {
					symbol = 1
				}
				allow_blank = false
				if zeros {
					continue
				}
				fallthrough
			case ch > 48 && ch <=57 :
				zeros = false
				allow_blank = false
				if symbol == 0 {
					symbol = 1
				}
				if start < 0 {
					start = i
				}
				end = i + 1
			default:
				goto OUT
		}
	}

OUT:

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

	var i64 int64 = 0

	for i, v := range str[start : end] {
		i64 += int64(v) - 48
		if i < end-start-1 {
			i64 *= 10
		}
	}

	i64 *= int64(symbol)

	if i64 < int64(min_32) {
		return min_32
	} else if i64 > int64(max_32) {
		return max_32
	}

	return int(i64)
}
