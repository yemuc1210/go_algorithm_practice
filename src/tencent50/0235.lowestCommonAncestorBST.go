package tencent50
/*简单
235. 二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，
	最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 
		x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/
/* 树的题目集中刷一次
递归
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p==nil || q==nil || root==nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val { //左侧
		return lowestCommonAncestor(root.Left,p,q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right,p,q)
	}
	return root  //q p在root两边
}