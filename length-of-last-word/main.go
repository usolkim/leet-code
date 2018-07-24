package main

import (
	"strings"
	"fmt"
)

func main(){
	fmt.Println(lengthOfLastWord("hello world"))
	fmt.Println(lengthOfLastWord("hello"))
	fmt.Println(lengthOfLastWord("hello world "))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("     "))
}

func lengthOfLastWord(s string) int {
	sTrim := strings.Trim(s, " ")
	if len(sTrim) == 0 {
		return 0
	}
	ss := strings.Split(sTrim, " ")
	l := len(ss)
	return len(ss[l-1])
}