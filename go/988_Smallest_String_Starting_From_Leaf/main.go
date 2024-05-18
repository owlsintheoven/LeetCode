package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	smallest := string(26 + 97)
	dfs(root, "", &smallest)
	return smallest
}

func dfs(node *TreeNode, path string, smallest *string) {
	if node == nil {
		return
	}
	curr := string(node.Val+97) + path
	if node.Left == nil && node.Right == nil {
		if *smallest > curr {
			*smallest = curr
		}
	}
	dfs(node.Left, curr, smallest)
	dfs(node.Right, curr, smallest)
}

func main() {
	//node3_1 := TreeNode{
	//	Val: 3,
	//}
	//node3_2 := TreeNode{
	//	Val: 4,
	//}
	//node3_3 := TreeNode{
	//	Val: 3,
	//}
	//node3_4 := TreeNode{
	//	Val: 4,
	//}
	//node2_1 := TreeNode{
	//	Val:   1,
	//	Left:  &node3_1,
	//	Right: &node3_2,
	//}
	//node2_2 := TreeNode{
	//	Val:   2,
	//	Left:  &node3_3,
	//	Right: &node3_4,
	//}
	//root := TreeNode{
	//	Val:   0,
	//	Left:  &node2_1,
	//	Right: &node2_2,
	//}
	//node3_1 := TreeNode{
	//	Val: 1,
	//}
	//node3_2 := TreeNode{
	//	Val: 3,
	//}
	//node3_3 := TreeNode{
	//	Val: 0,
	//}
	//node3_4 := TreeNode{
	//	Val: 2,
	//}
	//node2_1 := TreeNode{
	//	Val:   1,
	//	Left:  &node3_1,
	//	Right: &node3_2,
	//}
	//node2_2 := TreeNode{
	//	Val:   3,
	//	Left:  &node3_3,
	//	Right: &node3_4,
	//}
	//root := TreeNode{
	//	Val:   25,
	//	Left:  &node2_1,
	//	Right: &node2_2,
	//}
	node4_1 := TreeNode{
		Val: 0,
	}
	node3_1 := TreeNode{
		Val:  1,
		Left: &node4_1,
	}
	node3_2 := TreeNode{
		Val: 0,
	}
	node2_1 := TreeNode{
		Val:   2,
		Right: &node3_1,
	}
	node2_2 := TreeNode{
		Val:  1,
		Left: &node3_2,
	}
	root := TreeNode{
		Val:   2,
		Left:  &node2_1,
		Right: &node2_2,
	}
	fmt.Println(smallestFromLeaf(&root))
	//fmt.Println(dfs(&root))
}
