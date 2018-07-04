package main

import (
  "fmt"
)

func main() {
  fmt.Println(isPalindrome(121))
  fmt.Println(isPalindrome(-121))
  fmt.Println(isPalindrome(10))
  fmt.Println(isPalindrome(11))
  fmt.Println(isPalindrome(111))
  fmt.Println(isPalindrome(1111))
  fmt.Println(isPalindrome(1221))
  fmt.Println(isPalindrome(12321))
  fmt.Println(isPalindrome(123211))
  fmt.Println(isPalindrome(12331))
  fmt.Println(isPalindrome(1233))
  fmt.Println(isPalindrome(1212))
  fmt.Println(isPalindrome(121200))
  fmt.Println(isPalindrome(10001))
}

func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }
  if x < 10 {
    return true
  }

  var a [10]int
  var end int
  
  for i:=0;i<10;i++ {
    a[i] = x % 10
    x /=10
    if x == 0 {
      end = i + 1
      break
    }
  }

  for i:=0;i<end /2;i++ {
    if a[i] != a[end-i -1] {
      return false
    }
  }
  return true
}