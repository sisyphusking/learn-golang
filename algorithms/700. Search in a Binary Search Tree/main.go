package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)

}

func main() {
	l1 := &TreeNode{Val: 1}
	r1 := &TreeNode{Val: 3}

	l := &TreeNode{Val: 2, Left: l1, Right: r1}

	r := &TreeNode{Val: 7}

	root := &TreeNode{
		Val:   4,
		Left:  l,
		Right: r,
	}
	fmt.Println(searchBST(root, 2))
}
