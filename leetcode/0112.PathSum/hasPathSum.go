package lt112
import "go_practice/structs"
type TreeNode = structs.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。

叶子节点 是指没有子节点的节点。

*/


func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right==nil{
		return targetSum == root.Val
	}
	return hasPathSum(root.Left,targetSum-root.Val) || hasPathSum(root.Right,targetSum-root.Val)

}
