package offer54

import (

	"structs"
)
type TreeNode=structs.TreeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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