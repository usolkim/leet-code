package main

import "fmt"

func main() {
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("999", "999"))
	fmt.Println(multiply("9999242", "1321"))
	fmt.Println(multiply("9133", "0"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l := len(num1) + len(num2)
	nums := make([]int, l)

	for i := 0; i < len(num1); i++ {
		n1 := int(num1[i] - '0')
		for j := 0; j < len(num2); j++ {
			n2 := int(num2[j] - '0')
			nums[l-i-j-2] += n1 * n2
		}
	}

	s := make([]byte, l)
	for i := 0; i < l; i++ {
		if nums[i] > 9 {
			nums[i+1] += nums[i] / 10
			nums[i] %= 10
		}
		s[l-i-1] = byte(nums[i] + '0')
	}

	if s[0] == '0' {
		return string(s[1:])
	}

	return string(s)
}
