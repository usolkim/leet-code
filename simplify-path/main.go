package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath(""))
	fmt.Println(simplifyPath("////"))
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("home"))
	fmt.Println(simplifyPath("//home//"))
	fmt.Println(simplifyPath("./home/"))
	fmt.Println(simplifyPath("/./home/"))
	fmt.Println(simplifyPath("././home/"))
	fmt.Println(simplifyPath(".././home/"))
	fmt.Println(simplifyPath("/.././home/"))
	fmt.Println(simplifyPath("/a/./b/./c/"))
	fmt.Println(simplifyPath("/a/./b/../c/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a/./b/../../../../c/"))
	fmt.Println(simplifyPath("/a/./././././b/../../../../c/"))
	fmt.Println(simplifyPath("/a +_$^&b/"))
	fmt.Println(simplifyPath("/a/./b///../c/../././../d/..//../e/./f/./g/././//.//h///././/..///"))
}

func simplifyPath(path string) string {
	if path == "" {
		return "/"
	}
	paths := strings.Split(path, "/")
	tmp := make([]string, 0)
	i := 0
	for _, s := range paths {
		if s == "." || s == "" {
			continue
		} else if s == ".." {
			if i > 0 {
				i--
			}
			tmp = tmp[:i]
		} else {
			tmp = append(tmp, s)
			i = len(tmp)
		}
	}
	res := make([]byte, 0)
	for _, s := range tmp {
		if s == "" {
			continue
		}
		res = append(res, '/')
		res = append(res, s...)
	}
	if len(res) == 0 {
		return "/"
	}
	return string(res)
}
