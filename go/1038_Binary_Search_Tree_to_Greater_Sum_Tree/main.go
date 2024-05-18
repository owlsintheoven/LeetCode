package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	toSum(root, 0)
	return root
}

func toSum(node *TreeNode, prev int) int {
	if node == nil {
		return prev
	}
	sum := toSum(node.Right, prev)
	node.Val = node.Val + sum
	sum = toSum(node.Left, node.Val)
	return sum
}

func traverse(node *TreeNode) {
	nodes := []*TreeNode{node}
	var res []int
	for len(nodes) != 0 {
		var next []*TreeNode
		for _, n := range nodes {
			if n == nil {
				res = append(res, 0)
			} else {
				res = append(res, n.Val)
				next = append(next, n.Left)
				next = append(next, n.Right)
			}
		}
		nodes = next
	}
	fmt.Println(res)
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	traverse(root)
	res := bstToGst(root)
	traverse(res)
}
