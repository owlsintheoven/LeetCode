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

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	nodes := []*TreeNode{root}
	d := 1
	for {
		var next []*TreeNode
		if d == depth-1 {
			parents := nodes
			for _, p := range parents {
				if p != nil {
					leftTree := p.Left
					rightTree := p.Right
					p.Left = &TreeNode{
						Val:  val,
						Left: leftTree,
					}
					p.Right = &TreeNode{
						Val:   val,
						Left:  nil,
						Right: rightTree,
					}
				}
			}
			break
		}
		for _, n := range nodes {
			if n != nil {
				next = append(next, n.Left)
				next = append(next, n.Right)
			}
		}
		nodes = next
		d += 1
	}
	return root
}

func traverse(node *TreeNode) {
	nodes := []*TreeNode{node}
	var res []string
	for len(nodes) != 0 {
		var next []*TreeNode
		for _, n := range nodes {
			if n != nil {
				res = append(res, strconv.Itoa(n.Val))
				next = append(next, n.Left)
				next = append(next, n.Right)
			} else {
				res = append(res, "nil")
			}
		}
		nodes = next
	}
	fmt.Println(res)
}

func main() {
	//node3_1 := TreeNode{
	//	Val: 3,
	//}
	//node3_2 := TreeNode{
	//	Val: 1,
	//}
	//node3_3 := TreeNode{
	//	Val: 5,
	//}
	//node2_1 := TreeNode{
	//	Val:   2,
	//	Left:  &node3_1,
	//	Right: &node3_2,
	//}
	//node2_2 := TreeNode{
	//	Val:  6,
	//	Left: &node3_3,
	//}
	//root := TreeNode{
	//	Val:   4,
	//	Left:  &node2_1,
	//	Right: &node2_2,
	//}
	node3_1 := TreeNode{
		Val: 3,
	}
	node3_2 := TreeNode{
		Val: 1,
	}
	node2_1 := TreeNode{
		Val:   2,
		Left:  &node3_1,
		Right: &node3_2,
	}
	root := TreeNode{
		Val:  4,
		Left: &node2_1,
	}
	traverse(&root)
	root2 := addOneRow(&root, 1, 3)
	fmt.Println("root2")
	traverse(root2)
}
