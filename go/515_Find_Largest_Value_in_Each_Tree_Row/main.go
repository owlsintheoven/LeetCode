package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		m := nodes[0].Val
		var next []*TreeNode
		for _, n := range nodes {
			if n.Val > m {
				m = n.Val
			}
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		res = append(res, m)
		nodes = next
	}
	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	fmt.Println(largestValues(root))
}
