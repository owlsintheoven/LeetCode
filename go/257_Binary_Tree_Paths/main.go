package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root != nil {
		constructPaths(root, strconv.Itoa(root.Val), &res)
	}
	return res
}

func constructPaths(node *TreeNode, path string, res *[]string) {
	if node.Left == nil && node.Right == nil {
		*res = append(*res, path)
		return
	}
	if node.Left != nil {
		constructPaths(node.Left, path+"->"+strconv.Itoa(node.Left.Val), res)
	}
	if node.Right != nil {
		constructPaths(node.Right, path+"->"+strconv.Itoa(node.Right.Val), res)
	}
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(binaryTreePaths(root))
}
