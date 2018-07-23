package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println(totalNQueens1(11))
	elapsedTime := time.Since(startTime)
	fmt.Printf("totalNQueens1 실행시간: %s\n", elapsedTime)

	startTime = time.Now()
	fmt.Println(totalNQueens2(11))
	elapsedTime = time.Since(startTime)
	fmt.Printf("totalNQueens2 실행시간: %s\n", elapsedTime)

	startTime = time.Now()
	fmt.Println(totalNQueens3(11))
	elapsedTime = time.Since(startTime)
	fmt.Printf("totalNQueens3 실행시간: %s\n", elapsedTime)

	// fmt.Println(totalNQueens2(4))

	startTime = time.Now()
	// fmt.Println(totalNQueens(0))
	// fmt.Println(totalNQueens(1))
	// fmt.Println(totalNQueens(2))
	// fmt.Println(totalNQueens(3))
	// fmt.Println(totalNQueens(4))
	// fmt.Println(totalNQueens(5))
	// fmt.Println(totalNQueens(6))
	// fmt.Println(totalNQueens(7))
	// fmt.Println(totalNQueens(8))
	// fmt.Println(totalNQueens(9))
	// fmt.Println(totalNQueens(10))
	// fmt.Println(totalNQueens(11))
	// fmt.Println(totalNQueens(12))
	// fmt.Println(totalNQueens(13))
	fmt.Println(totalNQueens3(16))
	elapsedTime = time.Since(startTime)

	fmt.Printf("totalNQueens3 실행시간: %s\n", elapsedTime)

	// startTime = time.Now()
	// fmt.Println(totalNQueens2(16))
	// elapsedTime = time.Since(startTime)

	// fmt.Printf("totalNQueens2 실행시간: %s\n", elapsedTime)
}

// BackTracking. Time : O(N^2), Space : O(1)
func totalNQueens3(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	res := 0
	backt(0, 0, 0, int((1<<uint(n))-1), &res)
	return res
}

func backt(row, left, right, n int, res *int) {
	if row == n {
		*res++
		return
	}
	es := n & (^(row | left | right))
	for es > 0 {
		pos := es & (^es + 1)
		es -= pos
		backt(row|pos, (left|pos)<<1, (right|pos)>>1, n, res)
	}
}

// BackTracking. Time : O(N^2), Space : O(5N)
func totalNQueens2(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	res := 0
	cols := make([]int, n)
	lt := make([]int, (n-1)*2+1)
	rt := make([]int, (n-1)*2+1)
	bt(0, n, cols, lt, rt, &res)
	return res
}

func bt(r, n int, cols, lt, rt []int, ret *int) {
	for i := 0; i < n; i++ {

		if cols[i] == 1 || lt[r+i] == 1 || rt[n-1+r-i] == 1 {
			continue
		}
		if r+1 == n {
			*ret++
			return
		}
		cols[i] = 1
		lt[r+i] = 1
		rt[n-1+r-i] = 1

		bt(r+1, n, cols, lt, rt, ret)

		cols[i] = 0
		lt[r+i] = 0
		rt[n-1+r-i] = 0
	}
}

// BackTracking. Time : O(N^3), Space : O(N)
func totalNQueens1(n int) int {
	if n == 0 || n == 2 {
		return 0
	} else if n == 1 {
		return 1
	}
	q := make([]int, n+1)
	res := 0
	find(1, q, n, &res)

	return res

}

func find(r int, q []int, n int, res *int) {
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
			*res++
			return
		}
		find(r+1, q, n, res)
	}
	return
}
