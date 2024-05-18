package main

import "fmt"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// method1: dfs
func getImportance(employees []*Employee, id int) int {
	adj := make(map[int]*Employee)
	for _, e := range employees {
		adj[e.Id] = e
	}
	return dfs(id, adj)
}

func dfs(id int, adj map[int]*Employee) int {
	res := adj[id].Importance
	for _, s := range adj[id].Subordinates {
		res += dfs(s, adj)
	}
	return res
}

// method2: construct a tree and calculate
//type EmployeeNode struct {
//	Id         int
//	Importance int
//	Children   []*EmployeeNode
//}
//
//func getImportance(employees []*Employee, id int) int {
//	root := construct(employees)
//	// find
//	var boss *EmployeeNode
//	nodes := []*EmployeeNode{root}
//	for len(nodes) != 0 {
//		var next []*EmployeeNode
//		for _, n := range nodes {
//			if n.Id == id {
//				boss = n
//				break
//			}
//			next = append(next, n.Children...)
//		}
//		nodes = next
//	}
//	// sum
//	var sum int
//	nodes = []*EmployeeNode{boss}
//	for len(nodes) != 0 {
//		var next []*EmployeeNode
//		for _, n := range nodes {
//			sum += n.Importance
//			next = append(next, n.Children...)
//		}
//		nodes = next
//	}
//
//	return sum
//}
//
//func construct(employees []*Employee) *EmployeeNode {
//	// EmployeeNode array
//	nodes := make([]*EmployeeNode, 2000)
//	// parents[i] indicates the parent id of employees[i]
//	parents := make([]int, 2000)
//	for _, e := range employees {
//		t := &EmployeeNode{
//			Id:         e.Id,
//			Importance: e.Importance,
//		}
//		// connect children
//		for _, c := range e.Subordinates {
//			parents[c] = e.Id
//			if nodes[c] != nil {
//				t.Children = append(t.Children, nodes[c])
//			}
//		}
//		// connect parent
//		if parents[e.Id] != 0 {
//			nodes[parents[e.Id]].Children = append(nodes[parents[e.Id]].Children, t)
//		}
//		nodes[e.Id] = t
//	}
//	var root *EmployeeNode
//	for i, p := range parents {
//		if nodes[i] != nil && p == 0 {
//			root = nodes[i]
//			break
//		}
//	}
//	return root
//}
//
//func traverse(node *EmployeeNode) {
//	var res []int
//	nodes := []*EmployeeNode{node}
//	for len(nodes) != 0 {
//		var next []*EmployeeNode
//		for _, n := range nodes {
//			res = append(res, n.Id)
//			if len(n.Children) != 0 {
//				next = append(next, n.Children...)
//			}
//		}
//		nodes = next
//	}
//	fmt.Println(res)
//}

func main() {
	employees := []*Employee{
		{
			Id:           2,
			Importance:   3,
			Subordinates: nil,
		},
		{
			Id:           1,
			Importance:   5,
			Subordinates: []int{2, 3},
		},
		{
			Id:           3,
			Importance:   3,
			Subordinates: nil,
		},
	}
	fmt.Println(getImportance(employees, 1))
}
