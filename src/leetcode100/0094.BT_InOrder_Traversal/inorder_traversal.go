/*

二叉树的中序遍历
*/

package leetcode

import "structs"

type TreeNode = structs.TreeNode
// type TreeNode struct{
// 	Val int
// 	Left *TreeNode
// 	Right *TreeNode
// }
func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}

//迭代法 用到栈
func inorderTraversal_1(root *TreeNode)[]int{
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