package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")
	if len(pattern) != len(ss) {
		return false
	}
	ptow := make([]string, 256)
	wtop := make(map[string]uint8)
	for i := 0; i < len(pattern); i++ {
		if ptow[pattern[i]] == "" {
			ptow[pattern[i]] = ss[i]
		} else if ptow[pattern[i]] != ss[i] {
			return false
		}
		if wtop[ss[i]] == 0 {
			wtop[ss[i]] = pattern[i]
		} else if wtop[ss[i]] != pattern[i] {
			return false
		}
	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog dog dog dog"
	fmt.Println(wordPattern(pattern, s))
}
