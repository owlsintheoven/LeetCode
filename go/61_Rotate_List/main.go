package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	l := 0
	node := head
	var tail *ListNode
	for node != nil {
		l++
		tail = node
		node = node.Next
	}
	k = k % l
	if k == 0 {
		return head
	}
	var prev *ListNode
	newHead := head
	for i := 0; i < l-k; i++ {
		prev = newHead
		newHead = newHead.Next
	}
	tail.Next = head
	prev.Next = nil
	return newHead
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
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	k := 2
	traverse(head)
	res := rotateRight(head, k)
	traverse(res)
}
