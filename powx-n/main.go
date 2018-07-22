package main

import "fmt"

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
	fmt.Println(myPow(100.00000, 3))
	fmt.Println(myPow(-100.00000, 3))
	n := 1<<31 - 1
	fmt.Println(myPow(100, n))
}

func myPow(x float64, n int) (res float64) {
	if n == 0 {
		return 1.0
	}

	minus := false

	if n < 0 {
		n *= -1
		minus = true
	}

	res = pow(x, n)
	if minus {
		return 1.0 / res
	}
	return res
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n == 2 {
		return x * x
	}
	r := n % 2
	n /= 2
	p := pow(x, n)
	if r == 1 {
		return p * p * x
	}
	return p * p
}
