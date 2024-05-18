package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	var sign int
	if dividend*divisor > 0 {
		sign = 1
		if dividend < 0 {
			dividend *= -1
			divisor *= -1
		}
	} else {
		sign = -1
		if dividend < 0 {
			dividend *= -1
		} else {
			divisor *= -1
		}
	}
	res := 0
	for dividend >= divisor {
		dividend -= divisor
		res += 1
	}
	res *= sign
	return res
}

func main() {
	dividend := -2147483648
	divisor := -1
	fmt.Println(divide(dividend, divisor))
}
