package main

import (
	"fmt"
	"sort"
)

func minRemoveToMakeValid(s string) string {
	var leftIdx, removed []int
	leftC := 0
	for i, c := range s {
		if c == '(' {
			leftIdx = append(leftIdx, i)
			leftC += 1
		} else if c == ')' {
			if leftC == 0 {
				removed = append(removed, i)
			} else {
				leftC -= 1
				leftIdx = leftIdx[:len(leftIdx)-1]
			}
		}
	}
	if leftC > 0 {
		// remove left parenthesis
		removed = append(removed, leftIdx[:leftC]...)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(removed)))
	for _, i := range removed {
		s = s[:i] + s[i+1:]
	}
	return s
}

func main() {
	//s := "lee(t(c)o)de))"
	//s := "a)b(c)d"
	//s := "))(("
	s := "())()((("
	// (())) remove 1 )
	// ))((  remove all
	// ((()) remove 1 (
	// (((()))
	fmt.Println(minRemoveToMakeValid(s))
}
