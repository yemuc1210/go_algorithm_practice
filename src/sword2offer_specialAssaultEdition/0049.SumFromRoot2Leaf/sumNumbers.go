package offerS49

import "structs"
type TreeNode = structs.TreeNode
/*
剑指 Offer II 049. 从根节点到叶节点的路径数字之和
给定一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。

每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。

*/
func sumNumbers(root *TreeNode) int {
	sumRes := 0
	var dfs  func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		sum = sum*10 + root.Val
		if root.Left != nil {
			dfs(root.Left, sum)
		}
		if root.Right != nil {
			dfs(root.Right,sum)
		}
		if root.Left == nil && root.Right == nil {
			sumRes += sum
		}

	}
	dfs(root,0)
	return sumRes
}
