// manacher's algorithm
package main

import (
  "fmt"
  "strings"
)

func main(){
  fmt.Println(longestPalindrome("babad"))
  fmt.Println(longestPalindrome("cbbd"))
  fmt.Println(longestPalindrome("bananac"))
  fmt.Println(longestPalindrome("anana"))
  fmt.Println(longestPalindrome("bb"))
  fmt.Println(longestPalindrome("abb"))
}

func longestPalindrome(s string) string {
  var r, p, maxA, maxI int = -1, -1, -1, -1
  var A [2001]int

  var str_tmp string
  for _, ch := range s {
    str_tmp = str_tmp + "#" + string(ch)
  }
  str_tmp = str_tmp + "#"
  s = str_tmp
  var slen = len(s)

  for i := 0; i < slen; i++{
    if i > r {
      A[i] = 0
    }else{
      if p * 2 - i >= 0 && A[p * 2 - i] <= r - i {
        A[i] = A[p * 2 - i]
      } else {
        A[i] = r - i
      }
    }

    for i-A[i]-1 >= 0 && i+A[i]+1 < slen && s[i-A[i]-1] == s[i+A[i]+1] {
      A[i] = A[i] + 1
    }

    if r < A[i] + i {
      r = A[i] + i
      p = i
    }

    if maxA < A[i] {
      maxA = A[i]
      maxI = i
    }
  }

  return strings.Replace(s[maxI - maxA:maxI + maxA + 1], "#", "", -1)
}