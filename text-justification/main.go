package main

import "fmt"

func main() {
	print(fullJustify([]string{"ab", "ab", "ab", "Ab", "aa", "bb"}, 5))
	print(fullJustify([]string{"ab", "ab", "ab", "Ab", "a", "b"}, 5))
	print(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	print(fullJustify([]string{"This", "", "is", "an", "example", "of", "text", "justification."}, 16))
	print(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	print(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}

func print(words []string){
	for _,w :=range words{
		fmt.Println(w)
	}
	fmt.Println("----------------------------------")
}

func fullJustify(words []string, maxWidth int) []string {
	l := len(words)
	if l == 0 {
		return []string{}
	}
	res := make([]string, 0)
	blank := make([]int, l)

	cl := 0
	start := 0
	for i, word := range words {
		wl := len(word)
		if cl+wl+i-start > maxWidth {
			if i - start == 1 {
				blank[start] = maxWidth - cl
			}else{
				bc := (maxWidth - cl) / (i - start - 1)
				rc := (maxWidth - cl) % (i - start - 1)
				for j := start; j < i-1; j++ {
					blank[j] = bc
					if j < rc+start {
						blank[j]++
					}
				}
			}

			s := make([]byte, 0)
			for j := start; j < i; j++ {
				s = append(s, []byte(words[j])...)
				for k := 0; k < blank[j]; k++ {
					s = append(s, ' ')
				}
			}
			res = append(res, string(s))
			cl = 0
			start = i
		}
		cl += wl
	}
	s := make([]byte, 0)
	cl = 0
	for i:=start;i<l;i++{
		s = append(s, []byte(words[i])...)
		cl += len(words[i])
		if i == l-1 {
			for j:=0;j<maxWidth-cl;j++{
			s = append(s, ' ')
			}
		}else{
			cl++
			s = append(s, ' ')
		}
	}
	res = append(res, string(s))
	return res
}
