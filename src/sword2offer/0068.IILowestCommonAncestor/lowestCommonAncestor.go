package offer68

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 import "structs"
 type TreeNode = structs.TreeNode
 func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
	  return root
  }
  left := lowestCommonAncestor1(root.Left, p, q)
  right := lowestCommonAncestor1(root.Right, p, q)
  if left != nil {
	  if right != nil {
		  return root
	  }
	  return left
  }
  return right
}