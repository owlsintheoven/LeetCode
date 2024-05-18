package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var res int
	if root != nil {
		find(root, targetSum, []int{}, &res)
	}
	return res
}

func find(node *TreeNode, targetSum int, sums []int, res *int) {
	if node == nil {
		return
	}
	if node.Val == targetSum {
		*res++
	}
	for i, sum := range sums {
		if sum+node.Val == targetSum {
			*res++
		}
		sums[i] += node.Val
	}
	newSums := insert(sums, node.Val)
	find(node.Left, targetSum, newSums, res)
	find(node.Right, targetSum, newSums, res)
	for i := 0; i < len(sums); i++ {
		sums[i] -= node.Val
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
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: -2,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{
				Val: 11,
			},
		},
	}
	targetSum := 8
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//	},
	//}
	//targetSum := 2
	fmt.Println(pathSum(root, targetSum))
}
