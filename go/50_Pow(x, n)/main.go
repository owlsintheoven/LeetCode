package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	var res float64
	res = 1
	neg := false
	if n < 0 {
		n *= -1
		neg = true
	}
	for n > 0 {
		if n%2 == 0 {
			x *= x
			n /= 2
		} else {
			res *= x
			n--
		}
	}
	if neg {
		res = 1 / res
	}
	return res
}

//// powersOfTwo calculates ks where the sum of all 2 to k equals to n. n must be positive integer.
//func powersOfTwo(n int) []int {
//	var res []int
//	k := 0
//	p := 1
//	for n > p {
//		k++
//		p *= 2
//	}
//	if n == p {
//		res = append(res, k)
//	} else {
//		res = append(res, k-1)
//		res = append(res, powersOfTwo(n-p/2)...)
//	}
//	return res
//}

func main() {
	x := 2.00000
	n := 5
	fmt.Println(myPow(x, n))
}
