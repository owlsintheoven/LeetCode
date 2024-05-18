package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func traverse(node *ListNode) {
	var res []int
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	fmt.Println(res)
}

func doubleIt(head *ListNode) *ListNode {
	carry := double(head)
	if carry {
		next := head
		head = &ListNode{
			Val:  1,
			Next: next,
		}
	}
	return head
}

func double(node *ListNode) bool {
	if node == nil {
		return false
	}
	plusOne := double(node.Next)
	carry := false
	d := node.Val * 2
	if d >= 10 {
		node.Val = d - 10
		carry = true
	} else {
		node.Val = d
	}
	if plusOne {
		node.Val = node.Val + 1
	}
	return carry
}

func main() {
	node3 := &ListNode{
		Val:  9,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  9,
		Next: node3,
	}
	head := &ListNode{
		Val:  9,
		Next: node2,
	}
	traverse(head)
	head = doubleIt(head)
	traverse(head)
}
