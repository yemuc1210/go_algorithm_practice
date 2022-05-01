package interview04

/*
两个节点的第一个共同祖先
三种情况：
1 p 和 q 在 root的子树中，且分列 root 的 异侧（即分别在左、右子树中）；
2 p = root ，且 q 在 root 的左或右子树中；
3 q = root ，且 p 在 root 的左或右子树中；
*/
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil {
		if right != nil {   //说明p,q在root异侧
			return root
		}
		return left
	}
	// left和right同时为空  返回nil    right不为空，返回right
	return right
}