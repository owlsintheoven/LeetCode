package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	neg := false
	if targetSum < 0 {
		targetSum *= -1
		neg = true
	}
	if root != nil {
		findPath(root, targetSum, neg, []int{}, &res)
	}
	return res
}

func findPath(node *TreeNode, target int, neg bool, path []int, res *[][]int) {
	var newTarget int
	if !neg {
		newTarget = target - node.Val
	} else {
		newTarget = target + node.Val
	}
	if node.Left == nil && node.Right == nil {
		if newTarget == 0 {
			*res = append(*res, insert(path, node.Val))
		}
		return
	}
	newPath := insert(path, node.Val)
	if node.Left != nil {
		findPath(node.Left, newTarget, neg, newPath, res)
	}
	if node.Right != nil {
		findPath(node.Right, newTarget, neg, newPath, res)
	}
}

func insert(arr []int, n int) []int {
	var res []int
	res = append(res, arr...)
	res = append(res, n)
	return res
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(pathSum(root, 22))
}
