package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(1))
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(59))
	fmt.Println(intToRoman(60))
	fmt.Println(intToRoman(1883))
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(2444))
	fmt.Println(intToRoman(3999))
}

func intToRoman(num int) string {
	var roman []byte
	var r = [7]byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
	d := 1000
	for i := 0; i < 4; i++ {
		k := num / d
		num %= d
		d /= 10

		if k == 0 {
			continue
		}
		switch k % 5 {
		case 3:
			if k > 5 {
				roman = append(roman, r[i*2-1])
			}
			roman = append(roman, r[i*2], r[i*2], r[i*2])
		case 2:
			if k > 5 {
				roman = append(roman, r[i*2-1])
			}
			roman = append(roman, r[i*2], r[i*2])
		case 1:
			if k > 5 {
				roman = append(roman, r[i*2-1])
			}
			roman = append(roman, r[i*2])
		case 4:
			roman = append(roman, r[i*2], r[i*2-k/5-1])
		case 0:
			roman = append(roman, r[i*2-1])
		}
	}
	return string(roman)
}
