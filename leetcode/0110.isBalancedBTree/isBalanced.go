package lt110

import (
	"go_practice/structs"
)
type TreeNode=structs.TreeNode
/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

一般思路是层次遍历 dfs
*/
func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}
	if root.Left==nil && root.Right==nil{
		return true
	}
	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)
	return abs(leftHeight-rightHeight)<=1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}