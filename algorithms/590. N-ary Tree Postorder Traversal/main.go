package main

import (
	"container/list"
	"fmt"
)

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
	fmt.Println("postOrder : ", postOrder(root))
	fmt.Println("postorder : ", postorder(root))
	fmt.Println("preOrder : ", preOrder(root))
}

type Node struct {
	Val   int
	Nodes []*Node
}

//https://leetcode.com/problems/n-ary-tree-postorder-traversal/discuss/549341/Go%3A-Iterative-solution-(using-list)
//https://leetcode.com/problems/n-ary-tree-preorder-traversal/

//N叉树后序遍历
//使用递归的方式来解
func postOrder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	for _, node := range root.Nodes {
		res = append(res, postOrder(node)...)
	}

	res = append(res, root.Val)
	return res
}

func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := list.New()
	stack.PushFront(root)
	for stack.Len() > 0 {
		curr := stack.Remove(stack.Front()).(*Node)
		if curr != nil {
			res = append(res, curr.Val)
		}

		for i := 0; i < len(curr.Nodes); i++ {
			stack.PushFront(curr.Nodes[i])
		}
	}

	reverse(res)

	return res
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//N叉树前序遍历
//-------------------------------------------------

func preOrder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		//取出最后一个node
		r := stack[len(stack)-1]
		//取出stack中除最后一个剩余的node
		stack = stack[:len(stack)-1]

		res = append(res, r.Val)

		var tmp []*Node
		for _, v := range r.Nodes {
			//将取出的node倒序放到tmp中
			tmp = append([]*Node{v}, tmp...)
		}
		//将tmp放入到stack中
		stack = append(stack, tmp...)
	}
	return res
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	ret = append(ret, root.Val)
	for _, v := range root.Nodes {
		ret = append(ret, preorder(v)...)
	}
	return ret
}
