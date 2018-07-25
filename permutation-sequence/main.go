package main

import (
	"fmt"
)

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(1, 1))
	fmt.Println(getPermutation(2, 1))
	fmt.Println(getPermutation(2, 2))
	fmt.Println(getPermutation(3, 6))
	fmt.Println(getPermutation(4, 24))
	fmt.Println(getPermutation(7, 3653))
	fmt.Println(getPermutation(8, 12345))
	fmt.Println(getPermutation(8, 23746))
	fmt.Println(getPermutation(9, 2*3*4*5*6*7*8))
	fmt.Println(getPermutation(9, 2*3*4*5*6*7*8*9))
}

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	} else if n == 2 {
		if k == 1 {
			return "12"
		}
		return "21"
	}

	nums := make([]byte, n)
	mul := 1
	for i := 1; i <= n; i++ {
		nums[i-1] = byte('0' + i)
		mul *= i
	}

	return string(permutation(nums, k-1, n, mul))
}

func permutation(nums []byte, k, n, mul int) []byte {
	res := make([]byte, 0)
	if n == 3 {
		if k == 0 {
			res = append(res, nums[0], nums[1], nums[2])
		} else if k == 1 {
			res = append(res, nums[0], nums[2], nums[1])
		} else if k == 2 {
			res = append(res, nums[1], nums[0], nums[2])
		} else if k == 3 {
			res = append(res, nums[1], nums[2], nums[0])
		} else if k == 4 {
			res = append(res, nums[2], nums[0], nums[1])
		} else if k == 5 {
			res = append(res, nums[2], nums[1], nums[0])
		}
	} else {
		mul /= n
		a := k / mul
		b := k % mul
		res = append(res, nums[a])
		nnums := make([]byte, 0)
		nnums = append(nnums, nums[:a]...)
		nnums = append(nnums, nums[a+1:]...)

		res = append(res, permutation(nnums, b, n-1, mul)...)
	}
	return res
}
