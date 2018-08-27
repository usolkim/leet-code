package main

import "fmt"

func main(){
	fmt.Println(decodeAtIndex("leet2code3", 10))
	fmt.Println(decodeAtIndex("ha22", 5))
	fmt.Println(decodeAtIndex("a2345678999999999999999", 1))
}

func decodeAtIndex(S string, K int) string {
	l := len(S)
	var size uint64
	for i:=0;i<l;i++{
	  if isAlpha(S[i]) {
		size++
	  }else{
		size*=uint64(S[i] - '0')
	  }
	}
	for i:=l-1;i>=0;i--{
	  K = int(uint64(K) % size)
	  if K == 0 && isAlpha(S[i]) {
		return string(S[i])
	  }
	  if isAlpha(S[i]) {
		size--
	  }else{
		size /=uint64(S[i] - '0')
	  }
	}
	return "X"
  }
  
  func isAlpha(b byte) bool {
	return 'a' <= b && b <= 'z' 
  }