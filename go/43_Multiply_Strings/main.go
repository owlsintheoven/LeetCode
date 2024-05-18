package main

import (
	"fmt"
)

// `num1[i] * num2[j]` will be placed at indices `[i + j`, `i + j + 1]`
//
//	higher   lower
func multiply(num1 string, num2 string) string {
	digits := make([]uint8, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			m := (num1[i] - 48) * (num2[j] - 48)
			h, l := i+j, i+j+1
			s := m + digits[l]
			digits[h] += s / 10
			digits[l] = s % 10
		}
	}
	var res string
	for i := 0; i < len(digits); i++ {
		if len(res) != 0 || digits[i] != 0 {
			res = res + string(digits[i]+48)
		}
	}
	if len(res) == 0 {
		res = "0"
	}
	return res
}

// method: naive calculator
//func multiply(num1 string, num2 string) string {
//	res := "0"
//	u := 0
//	for i := len(num2) - 1; i >= 0; i-- {
//		a := multiplyByOneDigit(num1, num2[i]-48, u)
//		res = add(res, a)
//		u++
//	}
//	return res
//}
//
//func multiplyByOneDigit(num string, m uint8, unit int) string {
//	if num == "0" || m == 0 {
//		return "0"
//	}
//	var digits []uint8
//	for i := 0; i < unit; i++ {
//		digits = append(digits, 0)
//	}
//	var carry uint8
//	pos := len(num) - 1
//	for pos >= 0 {
//		n := num[pos] - 48
//		p := n * m
//		p += carry
//		r := p % 10
//		digits = append(digits, r)
//		carry = p / 10
//		pos--
//	}
//	if carry != 0 {
//		digits = append(digits, carry)
//	}
//	var res string
//	for _, d := range digits {
//		res = string(d+48) + res
//	}
//	return res
//}
//
//func add(num1, num2 string) string {
//	var digits []uint8
//	pos1 := len(num1) - 1
//	pos2 := len(num2) - 1
//	carry := false
//	for pos1 >= 0 || pos2 >= 0 {
//		var d1, d2 uint8
//		if pos1 < 0 {
//			d1 = 0
//		} else {
//			d1 = num1[pos1] - 48
//		}
//		if pos2 < 0 {
//			d2 = 0
//		} else {
//			d2 = num2[pos2] - 48
//		}
//		d := d1 + d2
//		if carry {
//			d++
//		}
//		if d >= 10 {
//			d -= 10
//			carry = true
//		} else {
//			carry = false
//		}
//		digits = append(digits, d)
//		pos1--
//		pos2--
//	}
//	if carry {
//		digits = append(digits, 1)
//	}
//	var res string
//	for _, d := range digits {
//		res = string(d+48) + res
//	}
//	return res
//}

func main() {
	//num1 := "123"
	//num2 := "456"
	num1 := "498828660196"
	num2 := "840477629533"
	//num1 := "9133"
	//num2 := "0"
	//num1 := "0"
	//num2 := "52"
	fmt.Println(multiply(num1, num2))
	//fmt.Println(add(num1, num2))
	//fmt.Println(multiplyByOneDigit("123", 6, 2))
}

// 000123
// 000456
// 000018
// 000138
