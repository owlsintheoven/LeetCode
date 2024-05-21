package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	var steps int
	exceed := dfs(root, &steps)
	if exceed < 0 {
		exceed *= -1
	}
	steps += exceed
	return steps
}

func dfs(node *TreeNode, steps *int) int {
	if node == nil {
		return 0
	}
	leftExceed := dfs(node.Left, steps)
	rightExceed := dfs(node.Right, steps)
	node.Val += leftExceed + rightExceed
	if leftExceed < 0 {
		leftExceed *= -1
	}
	if rightExceed < 0 {
		rightExceed *= -1
	}
	*steps += leftExceed + rightExceed
	return node.Val - 1
}

func main() {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	fmt.Println(distributeCoins(root))
}
