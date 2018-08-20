package main

import (
	"fmt"
)

func main() {
	fmt.Println(fractionToDecimal(-1, -2147483648))
	fmt.Println(fractionToDecimal(0, 0))
	fmt.Println(fractionToDecimal(1, 0))
	fmt.Println(fractionToDecimal(0, 15))
	fmt.Println(fractionToDecimal(10, -2))
	fmt.Println(fractionToDecimal(1, -3))
	fmt.Println(fractionToDecimal(1, 15))
	fmt.Println(fractionToDecimal(1, 12))
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(5, 3))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(2, 4))
	fmt.Println(fractionToDecimal(4, 2))
	fmt.Println(fractionToDecimal(4, 1))
	fmt.Println(fractionToDecimal(5, 7))
	fmt.Println(fractionToDecimal(2, 5))
	fmt.Println(fractionToDecimal(15, 7))
	fmt.Println(fractionToDecimal(51, 7))
	fmt.Println(fractionToDecimal(3, 11))
	fmt.Println(fractionToDecimal(1, 123442))

}

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return "NaN"
	} else if numerator == 0 {
		return "0"
	}
	var s string
	if numerator*denominator < 0 {
		s = "-"
	}
	if numerator < 0 {
		numerator *= -1
	}
	if denominator < 0 {
		denominator *= -1
	}
	a := make(map[int]int, 0)
	b := make([]int, 0)
	i := 1
	var recursive func(int) int
	recursive = func(r int) int {
		r *= 10
		if r == 0 {
			return -1
		}

		d := r % denominator
		// fmt.Println(i, r, d)
		if v, exist := a[d]; exist {
			if b[a[d]-1] != r/denominator {
				b = append(b, r/denominator)
				return v
			}
			return v - 1
		}
		b = append(b, r/denominator)
		a[d] = i
		i++
		return recursive(d)
	}
	c := numerator / denominator
	l := recursive(numerator % denominator)
	s = s + fmt.Sprintf("%d", c)
	if len(b) > 0 {
		s = s + "."
		for j, v := range b {
			if j == l {
				s = s + "("
			}
			s = s + fmt.Sprintf("%d", v)
		}
		if l >= 0 {
			s = s + ")"
		}
	}
	return s
}
