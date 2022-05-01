package lt653

import (
	"go_practice/structs"
	// "container/heap"
)
type TreeNode = structs.TreeNode
/*   lt653
剑指 Offer II 056. 二叉搜索树中两个节点之和
给定一个二叉搜索树的 根节点 root 和一个整数 k , 
请判断该二叉搜索树中是否存在两个节点它们的值之和等于 k 。
假设二叉搜索树中节点的值均唯一。

中序
*/
func findTarget(root *TreeNode, k int) bool {
	var find func(node *TreeNode, x int) bool
	find = func(node *TreeNode, x int) bool {
		if node == nil{
			return false
		}
		if x < node.Val {
			return find(node.Left, x)
		} else if x == node.Val {
			return true
		} else {
			return find(node.Right, x)
		}
	}
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node.Left != nil {
			if dfs(node.Left) {
				return true
			}
		}
		if node.Val*2 < k {
			if find(root, k-node.Val) {
				return true
			}
		} else {
			return false
		}
		if node.Right != nil {
			if dfs(node.Right) {
				return true
			}
		}
		return false
	}
	return dfs(root)
}

