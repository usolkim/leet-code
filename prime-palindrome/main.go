package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(primePalindrome(1))
	fmt.Println(primePalindrome(2))
	fmt.Println(primePalindrome(3))
	fmt.Println(primePalindrome(6))
	fmt.Println(primePalindrome(7))
	fmt.Println(primePalindrome(8))
	fmt.Println(primePalindrome(13))
	fmt.Println(primePalindrome(12623))
	fmt.Println(primePalindrome(45422541))

	fmt.Println(primePalindrome(100000000))
}

func primePalindrome(N int) int {
	if N == 1 {
		return 2
	}
	for i := N; i <= 2000000000; i++ {
		if 1000 <= i && i < 10000 {
			i = 10000
		} else if 100000 <= i && i < 1000000 {
			i = 1000000
		} else if 10000000 <= i && i < 100000000 {
			i = 100030001
		}
		if i == isPalindrome(i) && isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(N int) bool {
	if N == 1 {
		return false
	} else if N%2 == 0 {
		return N == 2
	} else if N%3 == 0 {
		return N == 3
	}
	s := int(math.Sqrt(float64(N)))
	for i := 5; i <= s; i += 2 {
		if N%i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(N int) int {
	s := 0
	for N > 0 {
		s = (10 * s) + (N % 10)
		N /= 10
	}
	return s
}
