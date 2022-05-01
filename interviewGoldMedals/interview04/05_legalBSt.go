package interview04

import "math"

/*
面试题 04.05. 合法二叉搜索树
实现一个函数，检查一棵二叉树是否为二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
func isValidBST(root *TreeNode) bool {
	return helper(root,math.MinInt64,math.MaxInt64)
}

func helper(root *TreeNode, lower,upper int) bool{
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val>= upper {
		return false  //数字上下界判断
	}

	return helper(root.Left,lower,root.Val) && helper(root.Right,root.Val,upper)
}