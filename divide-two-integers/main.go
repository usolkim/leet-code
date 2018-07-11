package main

import "fmt"

func main() {
	fmt.Println(divide(-2147483648, -1))
	fmt.Println(divide(-2147483648, 2))
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

	i := 0
	for dividend > divisor<<uint(i) {
		i++
	}

	ret := 0
	for dividend >= divisor {
		tmpDivisor := divisor << uint(i)
		if dividend >= tmpDivisor {
			dividend -= tmpDivisor
			ret += 1 << uint(i)
		}
		i--
	}

	ret *= sign
	max := 1<<31 - 1
	if ret > max {
		return max
	}
	return ret
}
