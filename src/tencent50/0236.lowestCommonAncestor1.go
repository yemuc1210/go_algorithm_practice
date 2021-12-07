package tencent50


/*
236. 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且
	 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	if left != nil {
			if right != nil {   //说明p,q在root异侧
					return root
			}
			return left
	}
	// left和right同时为空  返回nil    right不为空，返回right
	return right
}