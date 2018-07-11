package main

import "fmt"

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(10, 5))
	fmt.Println(divide(10, 1))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(-7, 3))
	fmt.Println(divide(-7, -3))
	fmt.Println(divide(10, 11))
	fmt.Println(divide(10, -11))
	fmt.Println(divide(-10, 11))
	fmt.Println(divide(-10, -11))
	fmt.Println(divide(-2147483648, 2))
}

func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	sign := 1
	if dividend < 0 {
		sign *= -1
		dividend *= -1
	}
	if divisor < 0 {
		sign *= -1
		divisor *= -1
	}
	// max := 1<<31 - 1
	divisorB := divisor
	for i := 0; i < 31; i++ {
		if dividend > divisorB {
			divisorB = divisorB << 1
		} else {
			break
		}
	}

	ret := 0
	for dividend >= divisor {
		if dividend >= divisorB {
			dividend -= divisorB
			ret++
			ret = ret << 1
		}
		divisorB = divisorB >> 1
	}
	// for dividend >= divisor {
	// 	dividend -= divisor
	// 	if ret < max {
	// 		ret++
	// 	}
	// }
	ret *= sign
	return ret
}
