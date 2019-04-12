package main

import "fmt"

func main() {

	t1 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	t2 := &TreeNode{
		Val:   3,
		Left:  t1,
		Right: nil,
	}
	t3 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	// tree 1
	t4 := &TreeNode{
		Val:   1,
		Left:  t2,
		Right: t3,
	}
	// *****************
	t5 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	t6 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	t7 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: t5,
	}
	t8 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: t6,
	}

	// tree 2
	t9 := &TreeNode{
		Val:   2,
		Left:  t7,
		Right: t8,
	}

	t := mergeTrees(t4, t9)

	fmt.Println(t)
	fmt.Println(t.Left.Val)
	fmt.Println(t.Right.Val)

}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}
