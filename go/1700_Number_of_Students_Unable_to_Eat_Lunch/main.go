package main

import "fmt"

//func countStudents(students []int, sandwiches []int) int {
//	for len(students) != 0 {
//		top := sandwiches[0]
//		first := -1
//		for i, preferred := range students {
//			if preferred == top {
//				first = i
//				break
//			}
//		}
//		if first == -1 {
//			return len(students)
//		}
//		students = append(students[first+1:], students[:first]...)
//		sandwiches = sandwiches[1:]
//	}
//	return 0
//}

func countStudents(students []int, sandwiches []int) int {
	counts := make([]int, 2)
	for _, preferred := range students {
		counts[preferred] += 1
	}
	remaining := len(students)
	for _, sandwich := range sandwiches {
		if counts[sandwich] == 0 {
			break
		}
		counts[sandwich] -= 1
		remaining -= 1
	}
	return remaining
}

func main() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	//students := []int{1, 1, 1, 0, 0, 1}
	//sandwiches := []int{1, 0, 0, 0, 1, 1}
	fmt.Println(countStudents(students, sandwiches))
}
