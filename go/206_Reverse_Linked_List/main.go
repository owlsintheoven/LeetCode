package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	n := head
	var prev *ListNode
	for n != nil {
		tmp := n.Next
		n.Next = prev
		prev = n
		n = tmp
	}
	return prev
}

func traverseList(head *ListNode) {
	var res []int
	n := head
	for n != nil {
		res = append(res, n.Val)
		n = n.Next
	}
	fmt.Println(res)
}

func main() {
	node5 := ListNode{
		Val:  5,
		Next: nil,
	}
	node4 := ListNode{
		Val:  4,
		Next: &node5,
	}
	node3 := ListNode{
		Val:  3,
		Next: &node4,
	}
	node2 := ListNode{
		Val:  2,
		Next: &node3,
	}
	head := ListNode{
		Val:  1,
		Next: &node2,
	}
	traverseList(&head)

	res := reverseList(&head)
	traverseList(res)
}
