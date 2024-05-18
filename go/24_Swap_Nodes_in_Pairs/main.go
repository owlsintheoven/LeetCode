package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	node := head
	var prev, first, second *ListNode
	for node != nil {
		if first == nil {
			first = node
			node = node.Next
		} else {
			second = node
			node = node.Next
			if prev != nil {
				prev.Next = second
			} else {
				head = second
			}
			second.Next = first
			first.Next = nil
			prev = first
			first = nil
			second = nil
		}
	}
	if first != nil && prev != nil {
		prev.Next = first
	}
	return head
}

//func swap(node *ListNode) {
//	if node == nil {
//		return nil, 0
//	}
//	next, order := swap(node.Next)
//	if order+1 == n {
//		return next, order + 1
//	}
//	node.Next = next
//	return node, order + 1
//}

func traverse(node *ListNode) {
	var res []int
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	fmt.Println(res)
}

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	head2 := swapPairs(&head)
	traverse(head2)
}
