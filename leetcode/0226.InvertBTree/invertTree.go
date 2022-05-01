package lt226
import "go_practice/structs"
type TreeNode = structs.TreeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left,root.Right = root.Right,root.Left
	return root
}