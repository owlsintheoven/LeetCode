package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var vals []int
	n := head
	for n != nil {
		vals = append(vals, n.Val)
		n = n.Next
	}
	n = head
	var prev *ListNode
	offset := 0
	for i := 1; i < 1+len(vals); i++ {
		if i >= left && i <= right {
			n2 := &ListNode{
				Val:  vals[right-1-offset],
				Next: nil,
			}
			offset += 1
			if prev != nil {
				prev.Next = n2
			}
			prev = n2
			if i == 1 {
				head = prev
			}
		} else {
			if prev != nil {
				prev.Next = n
			}
			prev = n
		}
		n = n.Next
	}
	return head
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

	res := reverseBetween(&head, 1, 5)
	traverseList(res)
}
