package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("0.1", "1.1"))
	fmt.Println(compareVersion("1.0.1", "1"))
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
}

func compareVersion(version1 string, version2 string) int {
	if version1 == version2 {
		return 0
	}
	if len(version1) == 0 || len(version2) == 0 {
		return 0
	}
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")
	l1 := len(s1)
	l2 := len(s2)

	l := l1
	if l < l2 {
		l = l2
	}

	for i := 0; i < l; i++ {
		v1, v2 := 0, 0
		if i < l1 {
			v1, _ = strconv.Atoi(s1[i])
		}
		if i < l2 {
			v2, _ = strconv.Atoi(s2[i])
		}
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	return 0
}
