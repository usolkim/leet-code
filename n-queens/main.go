package main

import (
	"fmt"
)

func main() {
	strs := solveNQueens(4)
	for _, ss := range strs {
		for _, s := range ss {
			fmt.Println(s)
		}
		fmt.Println()
	}

	strs = solveNQueens(5)
	for _, ss := range strs {
		for _, s := range ss {
			fmt.Println(s)
		}
		fmt.Println()
	}

	strs = solveNQueens(8)
	for _, ss := range strs {
		for _, s := range ss {
			fmt.Println(s)
		}
		fmt.Println()
	}
}

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{[]string{}}
	} else if n == 1 {
		return [][]string{[]string{"Q"}}
	} else if n == 2 {
		return [][]string{}
	}

	q := make([]int, n+1)
	qs := make([][]int, 0)
	res := make([][]string, 0)
	find(1, q, n, &qs)

	org := make([]byte, n)

	for i := 0; i < n; i++ {
		org[i] = '.'
	}

	for _, r := range qs {
		ss := make([]string, 0)
		for i := 1; i <= n; i++ {
			tmp := make([]byte, n)
			copy(tmp, org)
			tmp[r[i]-1] = 'Q'
			s := string(tmp)
			ss = append(ss, s)
		}
		res = append(res, ss)
	}

	return res

}

func find(r int, q []int, n int, res *[][]int) {
	for i := 1; i <= n; i++ {
		e := false
		for j := 1; j < r; j++ {
			if i == q[j] || q[j]-i == r-j || i-q[j] == r-j {
				e = true
				break
			}
		}
		if e == true {
			continue
		}
		q[r] = i
		if r == n {
			qq := make([]int, n+1)
			copy(qq, q)
			*res = append(*res, qq)
			return
		}
		find(r+1, q, n, res)
	}
	return
}
