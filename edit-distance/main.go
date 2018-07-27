package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDistance("", ""))
	fmt.Println(minDistance("a", "a"))
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("ros", "horse"))
	fmt.Println(minDistance("intention", "execution"))
	fmt.Println(minDistance("execution", "intention"))
	fmt.Println(minDistance("dinitrophenylhydrazine", "benzalphenylhydrazone"))
}

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	mem := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		mem[i] = make([]int, l2+1)
		for j := 0; j <= l2; j++ {
			mem[i][j] = -1
		}
	}
	return dp([]byte(word1), []byte(word2), mem)
}

func dp(w1, w2 []byte, mem [][]int) int {
	l1 := len(w1)
	l2 := len(w2)

	if l1 == 0 {
		mem[l1][l2] = l2
	} else if l2 == 0 {
		mem[l1][l2] = l1
	}

	if mem[l1][l2] != -1 {
		return mem[l1][l2]
	}

	if w1[l1-1] == w2[l2-1] {
		rs := dp(w1[:l1-1], w2[:l2-1], mem)
		mem[l1][l2] = rs
	} else {
		rm := dp(w1[:l1-1], w2[:l2], mem) + 1
		mem[l1][l2] = rm
		rp := dp(w1[:l1-1], w2[:l2-1], mem) + 1
		if rp < mem[l1][l2] {
			mem[l1][l2] = rp
		}
		in := dp(w1[:l1], w2[:l2-1], mem) + 1
		if in < mem[l1][l2] {
			mem[l1][l2] = in
		}
	}
	return mem[l1][l2]
}
