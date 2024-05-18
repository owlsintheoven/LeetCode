package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	return rle(n)
}

func rle(num int) string {
	if num == 1 {
		return "1"
	}
	s := rle(num - 1)
	var res string
	prev := int32(s[0])
	count := 1
	for _, c := range s[1:] {
		if c == prev {
			count++
		} else {
			res += strconv.Itoa(count)
			res += string(prev)
			prev = c
			count = 1
		}
	}
	res += strconv.Itoa(count)
	res += string(prev)
	return res
}

func main() {
	n := 4
	fmt.Println(countAndSay(n))
}
