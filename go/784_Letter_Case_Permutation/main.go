package main

import "fmt"

func letterCasePermutation(s string) []string {
	res := []string{""}
	for _, c := range s {
		var next []string
		if c >= 65 && c <= 90 {
			// upper case
			for _, r := range res {
				next = append(next, r+string(c))
				next = append(next, r+string(c+32))
			}
		} else if c >= 97 && c <= 122 {
			// lower case
			for _, r := range res {
				next = append(next, r+string(c))
				next = append(next, r+string(c-32))
			}
		} else {
			// digits
			for _, r := range res {
				next = append(next, r+string(c))
			}
		}
		res = next
	}
	return res
}

func main() {
	s := "a1b2"
	fmt.Println(letterCasePermutation(s))
}
