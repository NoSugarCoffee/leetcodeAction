package tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//104. 二叉树的最大深度
//给定一个二叉树，找出其最大深度。
//给定二叉树 [3,9,20,null,null,15,7]，返回它的最大深度 3 。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left:= maxDepth(root.Left) + 1
	right:=maxDepth(root.Right) + 1
	if  left > right {
		return left
	}else {
		return right
	}
}
