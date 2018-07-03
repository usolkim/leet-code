package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1563847412))
	fmt.Println(reverse(-1563847412))
	fmt.Println(reverse(2147483647))
	fmt.Println(reverse(-2147483648))
}

func reverse(x int) int {
	var minus bool
	if x < 0 {
		minus = true
		x *= -1
	}

	var r int
	var res int64 = 0
	for x > 0 {
		r = x % 10
		x /= 10
		res = res*10 + int64(r)
	}

	if minus {
		res *= -1
	}

	var max_32 int64 = 1<<31 - 1
	var min_32 int64 = -1 << 31

	if res < min_32 || res > max_32 {
		return 0
	}

	return int(res)
}
