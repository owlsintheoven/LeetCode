package main

import "fmt"

func lengthOfLastWord(s string) int {
	count := 0
	var lengths []int
	for _, c := range s {
		if c == ' ' {
			if count != 0 {
				lengths = append(lengths, count)
				count = 0
			}
		} else {
			count++
		}
	}
	if count != 0 {
		return count
	}
	return lengths[len(lengths)-1]
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}
