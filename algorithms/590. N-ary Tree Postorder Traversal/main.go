package main

import "fmt"

func main() {
	l1 := &Node{Val: 5}
	r1 := &Node{Val: 6}
	l := &Node{
		Val:   3,
		Nodes: []*Node{l1, r1},
	}
	m := &Node{Val: 2}
	r := &Node{Val: 4}

	root := &Node{
		Val:   1,
		Nodes: []*Node{l, m, r},
	}
	postOrder(root)
	fmt.Println(list)
}

type Node struct {
	Val   int
	Nodes []*Node
}

var list []int

func postOrder(root *Node) {
	if root == nil {
		return
	}
	for _, i := range root.Nodes {
		postOrder(i)
	}
	list = append(list, root.Val)
}
