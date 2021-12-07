package interview04

/*
面试题 04.04. 检查平衡性
实现一个函数，检查二叉树是否平衡。
在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。


计算每个节点左右子树的深度
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