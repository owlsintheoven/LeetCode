package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	// find first digit bigger than the next one
	for i := 0; i < k; i++ {
		biggerDigit := num[0]
		removeIdx := -1
		for j := 1; j < len(num); j++ {
			n := num[j]
			if biggerDigit > n {
				removeIdx = j
				break
			}
			biggerDigit = n
		}
		if removeIdx == -1 {
			num = num[:len(num)-1]
		} else {
			num = num[:removeIdx-1] + num[removeIdx:]
		}
	}
	// remove zero prefix
	nonZeroIdx := -1
	for i := 0; i < len(num); i++ {
		if num[i] != 48 {
			nonZeroIdx = i
			break
		}
	}
	if nonZeroIdx == -1 {
		return "0"
	}
	num = num[nonZeroIdx:]
	return num
}

func main() {
	//num := "1432219"
	//k := 3
	//num := "100000"
	//k := 1
	num := "12345"
	k := 1
	//num := "10"
	//k := 2
	fmt.Println(removeKdigits(num, k))
}

// 1432219
//  4    9
// 132219
//  3   9
// 12219
//  2  9
// 1219

// 14532219
//  45    9
// 123
//  23

// 122319
//  2 3 9
