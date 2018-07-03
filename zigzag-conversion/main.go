package main

import (
  "fmt"
  "strings"
)

func main(){
  fmt.Println(convert("PAYPALISHIRING", 3))
  fmt.Println(convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {
    if numRows == 1 {
      return s
    }
    ss := []string{}
    for i:=1; i<=numRows; i++ {
      k := i-1
      ss = append(ss, string(s[k]))
      for {
        k = k + 2*(numRows-i) - 1
        if k >= len(s) {
          break;
        }
        fmt.Println(numRows, i, k, len(s))
        ss = append(ss, string(s[k]))
      }
    }
    return strings.Join(ss, "")
}