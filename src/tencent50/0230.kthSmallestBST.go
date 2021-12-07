package tencent50
/* 中
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
*/
//中序遍历
//迭代法 用到栈
func inorderTraversal(root *TreeNode)[]int{
	res := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
	
		root = root.Right
	}
	return res
}
func kthSmallest(root *TreeNode, k int) int {
	order := inorderTraversal(root)
	n := len(order)
	if k>n{
		return -1
	}
	return order[k-1]
}