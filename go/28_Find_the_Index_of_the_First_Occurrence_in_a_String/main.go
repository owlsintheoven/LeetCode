package main

import "fmt"

func strStr(haystack string, needle string) int {
	p := 0
	for p < len(haystack) {
		if haystack[p] != needle[0] {
			p += 1
		} else {
			found := true
			for i := 1; i < len(needle); i++ {
				if p+i == len(haystack) || haystack[p+i] != needle[i] {
					found = false
					break
				}
			}
			if found {
				return p
			}
			p += 1
		}
	}
	return -1
}

func main() {
	//haystack := "sadbutsad"
	//needle := "sad"
	//haystack := "leetcode"
	//needle := "leeto"
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStr(haystack, needle))
}
