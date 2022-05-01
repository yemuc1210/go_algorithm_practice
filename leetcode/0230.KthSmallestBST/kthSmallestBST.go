package lt230
import "go_practice/structs"
type TreeNode = structs.TreeNode
//剑指Offer 54
//BST的中序遍历好像是
func kthSmallest(root *TreeNode, k int) int {
	order := inorderTraversal(root)
	n := len(order)
	if k>n{
		return -1
	}
	return order[k-1]
}
/*
给定一棵二叉搜索树，请找出其中第k大的节点。
二叉搜索树的中序序列是递增的序列
*/
func kthLargest(root *TreeNode, k int) int {
	order := inorderTraversal(root)
	// fmt.Printf("order:%v\n",order)
	n := len(order)
	//条件1<=k<=n   所以不用判断了
	return order[n-k]
}

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