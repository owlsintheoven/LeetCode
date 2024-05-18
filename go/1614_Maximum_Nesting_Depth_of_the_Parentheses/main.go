package main

import "fmt"

func maxDepth(s string) int {
	leftC, maxD := 0, 0
	for _, c := range s {
		if c == '(' {
			leftC += 1
			if maxD < leftC {
				maxD = leftC
			}
		} else if c == ')' {
			leftC -= 1
		}
	}
	return maxD
}

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	fmt.Println(maxDepth(s))
}
