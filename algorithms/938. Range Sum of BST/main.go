package main

import "fmt"

func main() {

	a := &TreeNode{
		Val: 7,
	}
	b := &TreeNode{
		Val: 5,
	}

	c := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: a,
	}

	d := &TreeNode{
		Val:   3,
		Left:  b,
		Right: c,
	}
	e := &TreeNode{
		Val:   7,
		Left:  d,
		Right: c,
	}

	fmt.Println(rangeSumBST1(e, 7, 15))
	//fmt.Println(rangeSumBST(e, 7, 15))

}

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {

	sum := 0

	if root != nil {
		if L <= root.Val && root.Val <= R {
			sum += root.Val
		}
		sum += rangeSumBST(root.Left, L, R)
		sum += rangeSumBST(root.Right, L, R)
	}

	return sum
}

// 同样的写法比上面那种快20ms
func rangeSumBST1(root *TreeNode, L int, R int) int {
	sum := 0
	if root != nil {
		if root.Val >= L && root.Val <= R {
			sum += root.Val
		}
	}
	if root.Left != nil {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Right != nil {
		sum += rangeSumBST(root.Right, L, R)
	}
	return sum
}
