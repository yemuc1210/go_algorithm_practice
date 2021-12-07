/*

二叉树的先序遍历
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
	preorder(root, &result)
	return result
}

func preorder(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right, output)
	}
}