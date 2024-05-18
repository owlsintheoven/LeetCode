package main

import "fmt"

func canJump(nums []int) bool {
	stayed := make([]bool, len(nums))
	stayed[0] = true
	next := []int{0}
	for len(next) != 0 {
		var tmp []int
		for _, n := range next {
			for i := 0; i <= nums[n]; i++ {
				if n+i == len(nums)-1 {
					return true
				}
				if !stayed[n+i] {
					tmp = append(tmp, n+i)
					stayed[n+i] = true
				}
			}
		}
		next = tmp
	}
	return false
}

func main() {
	//nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	//nums := []int{2, 0, 0}
	nums := []int{0}
	fmt.Println(canJump(nums))
}
