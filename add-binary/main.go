package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("101", "110"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("1", "1"))
	fmt.Println(addBinary("10", "10"))
	fmt.Println(addBinary("11", "11"))
	fmt.Println(addBinary("1111111", "11"))
	fmt.Println(addBinary("10101010101010", "1011"))
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	if la < lb {
		a, b = b, a
		la, lb = lb, la
	}

	res := make([]byte, la+1)
	res[0] = '0'

	for i, c := range a {
		res[i+1] = byte(c)
	}

	for i := lb - 1; i >= 0; i-- {
		res[i-lb+1+la] += b[i] - '0'
	}

	for i := la; i > 0; i-- {
		if res[i] > '1' {
			res[i] -= 2
			res[i-1]++
		}
	}

	if res[0] == '0' {
		return string(res[1:])
	}
	return string(res)
}
