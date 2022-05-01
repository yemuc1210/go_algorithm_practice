/*

二叉树的后序遍历
*/

package leetcode

import "go_practice/structs"

type TreeNode = structs.TreeNode

// type TreeNode struct{
// 	Val int
// 	Left *TreeNode
// 	Right *TreeNode
// }
func postorderTraversal(root *TreeNode) []int {
	var result []int
	postorder(root, &result)
	return result
}

func postorder(root *TreeNode, output *[]int) {
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		*output = append(*output, root.Val)
	}
}
