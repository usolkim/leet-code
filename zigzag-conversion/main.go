package main

import (
  "fmt"
)

func main() {
  fmt.Println(convert("PAYPALISHIRING", 3))
  fmt.Println(convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {
  if numRows == 1 {
    return s
  }
  s_len := len(s)
  res := make([]byte, s_len)
  k := 2 * (numRows - 1)
  ii := 0
  for i := 0; i < numRows && i < s_len; i++ {

    if i == 0 || i == numRows-1 {
      for j := i; j < s_len; j += k {
        res[ii] = s[j]
        ii++
      }
    } else {
      res[ii] = s[i]
      ii++
      l := k - (2 * i)
      for j := i + l; j < s_len; j += l {
        res[ii] = s[j]
        ii++
        l = k - l
      }
    }
  }
  return string(res)
}
