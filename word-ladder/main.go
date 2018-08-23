package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "doa", "lot", "cog"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hit", "dot", "doa", "lot", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	hasEndWord := false
	for _, w := range wordList {
		if w == endWord {
			hasEndWord = true
			break
		}
	}
	if !hasEndWord {
		return 0
	}

	queue := make([]string, 0)

	findWords(beginWord, wordList, &queue)

	cnt := 1
	for len(queue) > 0 {
		cnt++
		size := len(queue)
		for i := 0; i < size; i++ {
			nextWord := queue[0]
			queue = queue[1:]
			if nextWord == endWord {
				return cnt
			}
			findWords(nextWord, wordList, &queue)
		}
	}

	return 0
}

func findWords(word string, wordList []string, queue *[]string) {
	for i, w := range wordList {
		if w == "X" {
			continue
		}
		d := 0
		for j := 0; j < len(word); j++ {
			if word[j] != w[j] {
				d++
			}
			if d > 1 {
				break
			}
		}
		if d == 1 {
			*queue = append(*queue, w)
			wordList[i] = "X"
		}
	}
}
