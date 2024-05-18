package main

import (
	"fmt"
	"sort"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {
	var fractions []float64
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			fractions = append(fractions, float64(arr[i])/float64(arr[j]))
		}
	}
	sort.Float64s(fractions)
	target := fractions[k-1]
	var res []int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if target == float64(arr[i])/float64(arr[j]) {
				res = []int{arr[i], arr[j]}
			}
		}
	}
	return res
}

func main() {
	arr := []int{1, 2, 3, 5}
	k := 3
	fmt.Println(kthSmallestPrimeFraction(arr, k))
}
