package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
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
	node3 := &ListNode{
		Val:  9,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  1,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  5,
		Next: node2,
	}
	head := &ListNode{
		Val:  4,
		Next: node1,
	}
	traverse(head)
	deleteNode(head)
	traverse(head)
}
