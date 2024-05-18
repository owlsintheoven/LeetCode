package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head2, _ := remove(head, n)
	return head2
}

func remove(node *ListNode, n int) (*ListNode, int) {
	if node == nil {
		return nil, 0
	}
	next, order := remove(node.Next, n)
	if order+1 == n {
		return next, order + 1
	}
	node.Next = next
	return node, order + 1
}

func traverse(node *ListNode) {
	var res []int
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	fmt.Println(res)
}

func main() {
	//head := ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val: 5,
	//				},
	//			},
	//		},
	//	},
	//}
	//n := 2
	head := ListNode{
		Val: 1,
	}
	n := 1
	head2 := removeNthFromEnd(&head, n)
	traverse(head2)
}
