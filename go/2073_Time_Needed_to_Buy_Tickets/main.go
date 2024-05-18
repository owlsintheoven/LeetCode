package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	time := 0
	for {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] > 0 {
				tickets[i] -= 1
				time += 1
				if i == k && tickets[k] == 0 {
					return time
				}
			}
		}
	}
}

func main() {
	//tickets := []int{2, 3, 2}
	//k := 2
	tickets := []int{5, 1, 1, 1}
	k := 0
	fmt.Println(timeRequiredToBuy(tickets, k))
}
