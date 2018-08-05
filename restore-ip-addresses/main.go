package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("255255255255"))
	fmt.Println(restoreIpAddresses("1111"))
	fmt.Println(restoreIpAddresses("2552552550"))
	fmt.Println(restoreIpAddresses("25502550"))
	fmt.Println(restoreIpAddresses("22222222222"))
	fmt.Println(restoreIpAddresses("2222222222"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("9999"))
	fmt.Println(restoreIpAddresses("999"))
	fmt.Println(restoreIpAddresses("25625625562"))
	fmt.Println(restoreIpAddresses("256255255255"))
	fmt.Println(restoreIpAddresses("00000"))
}

func restoreIpAddresses(s string) []string {
	l := len(s)
	if l < 4 || l > 12 {
		return nil
	}

	res := make([]string, 0)
	ad := make([]string, 4)

	var dp func(c, start int)
	dp = func(c, start int) {
		if c == 4 && start == l {
			tmp := make([]string, 4)
			copy(tmp, ad)
			str := ""
			for i := 0; i < 4; i++ {
				str = str + tmp[i]
				if i < 3 {
					str = str + "."
				}
			}
			res = append(res, str)
		} else if c < 4 && start < l {
			if s[start] == '0' {
				ad[c] = "0"
				dp(c+1, start+1)
			} else {
				for i := 1; i <= 3 && start+i <= l; i++ {
					ipn := s[start : start+i]
					if i == 3 {
						if n, err := strconv.Atoi(ipn); err != nil || n > 255 {
							continue
						}
					}
					ad[c] = ipn
					dp(c+1, start+i)
				}
			}
		}
	}
	dp(0, 0)
	return res
}
