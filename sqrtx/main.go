package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(3))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(5))
	fmt.Println(mySqrt(6))
	fmt.Println(mySqrt(7))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(11))
	fmt.Println(mySqrt(12))
	fmt.Println(mySqrt(13))
	fmt.Println(mySqrt(14))
	fmt.Println(mySqrt(15))
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(163142352))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	} else if x < 4 {
		return 1
	} else if x < 9 {
		return 2
	}
	return bsearch(0, x/2, x)
}

func bsearch(s, e, x int) int {
	if s+1 == e {
		return s
	}
	i := (s + e) / 2
	j := i * i
	if j == x {
		return i
	} else if j < x {
		return bsearch(i, e, x)
	}
	return bsearch(s, i, x)
}
