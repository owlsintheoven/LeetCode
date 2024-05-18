package main

import (
	"fmt"
)

func trap(height []int) int {
	// In each iteration, if height is going up, compare the tallest height before the current point and get min of them
	// Use min as the limit to trap the water, iterate back to the tallest point and modify height with trapped water
	// Break iteration if the point with taller or same height as min is reached
	water := 0
	prevHeight := 0
	prevTallestHeight := -1
	height = append([]int{-1}, height...)
	for i := 0; i < len(height); i++ {
		if height[i] > prevHeight {
			var trapHeight int
			if height[i] < prevTallestHeight {
				trapHeight = height[i]
			} else {
				trapHeight = prevTallestHeight
				prevTallestHeight = height[i]
			}
			for j := i - 1; j >= 0; j-- {
				if height[j] >= trapHeight {
					break
				}
				water += trapHeight - height[j]
				height[j] = trapHeight
			}
		}
		prevHeight = height[i]
	}
	return water
}

func main() {
	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}
