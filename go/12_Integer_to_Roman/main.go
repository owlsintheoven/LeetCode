package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	roman := []string{"I", "V", "X", "L", "C", "D", "M"}
	var res string
	d := 0
	for num != 0 {
		n := num % 10
		num /= 10
		var s string
		if n >= 1 && n <= 3 {
			s = strings.Repeat(roman[d], n)
		} else if n == 4 {
			s = roman[d] + roman[d+1]
		} else if n == 5 {
			s = roman[d+1]
		} else if n >= 6 && n <= 8 {
			s = roman[d+1] + strings.Repeat(roman[d], n-5)
		} else if n == 9 {
			s = roman[d] + roman[d+2]
		}
		res = s + res
		d += 2
	}
	return res
}

func main() {
	num := 3
	fmt.Println(intToRoman(num))
}
