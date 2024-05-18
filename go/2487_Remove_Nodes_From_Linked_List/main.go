package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// method: reverse ->
func removeNodes(head *ListNode) *ListNode {
	// reverse
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr = prev
	head = curr
	prev = nil
	for curr != nil {
		next := curr.Next
		if prev != nil {
			if curr.Val >= prev.Val {
				prev.Next = curr
				prev = curr
			}
		} else {
			prev = curr
		}
		curr = next
	}
	prev.Next = nil
	// reverse
	curr = head
	prev = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr = prev
	return curr
}

// method: recursive
//func removeNodes(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	node := head
//	nextGreaterNode := removeNodes(node.Next)
//	node.Next = nextGreaterNode
//	if nextGreaterNode == nil || node.Val >= nextGreaterNode.Val {
//		return node
//	}
//	return nextGreaterNode
//}

// method: stack
//func removeNodes(head *ListNode) *ListNode {
//	var stack []int
//	node := head
//	for node != nil {
//		for len(stack) != 0 && stack[0] < node.Val {
//			stack = pop(stack)
//		}
//		stack = push(stack, node.Val)
//		node = node.Next
//	}
//	var curr, next *ListNode
//	for _, s := range stack {
//		curr = &ListNode{
//			Val:  s,
//			Next: nil,
//		}
//		curr.Next = next
//		next = curr
//	}
//	return curr
//}
//func push(stack []int, t int) []int {
//	res := []int{t}
//	res = append(res, stack...)
//	return res
//}
//
//func pop(stack []int) []int {
//	var res []int
//	res = append(res, stack[1:]...)
//	return res
//}

// naive method
//func removeNodes(head *ListNode) *ListNode {
//	var nodes []*ListNode
//	node := head
//	for node != nil {
//		var removed []int
//		for i, n := range nodes {
//			if n.Val < node.Val {
//				removed = append(removed, i)
//			}
//		}
//		for i := len(removed) - 1; i >= 0; i-- {
//			nodes = removeIdx(nodes, removed[i])
//		}
//		nodes = append(nodes, node)
//		node = node.Next
//	}
//	nodes[0].Next = nil
//	prev := nodes[0]
//	for _, n := range nodes[1:] {
//		prev.Next = n
//		prev = n
//	}
//	return nodes[0]
//}
//
//func removeIdx(nodes []*ListNode, idx int) []*ListNode {
//	var res []*ListNode
//	res = append(res, nodes[:idx]...)
//	res = append(res, nodes[idx+1:]...)
//	return res
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
	// 5 2 13 3 8
	node5 := &ListNode{
		Val:  8,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  3,
		Next: node5,
	}
	node3 := &ListNode{
		Val:  13,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}
	head := &ListNode{
		Val:  5,
		Next: node2,
	}
	traverse(head)
	head = removeNodes(head)
	traverse(head)
}
