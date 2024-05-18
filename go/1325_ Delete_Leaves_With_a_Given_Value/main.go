package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
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
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	traverse(root)
	res := removeLeafNodes(root, 2)
	traverse(res)
}
