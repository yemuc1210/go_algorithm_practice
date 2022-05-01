package lt236
import "go_practice/structs"
type TreeNode = structs.TreeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
/*
根据以上定义，若 root 是 p, q 的 最近公共祖先 ，则只可能为以下情况之一：

1 p 和 q 在 root的子树中，且分列 root 的 异侧（即分别在左、右子树中）；
2 p = root ，且 q 在 root 的左或右子树中；
3 q = root ，且 p 在 root 的左或右子树中；


先序遍历，需要p,q返回，当p,q在节点root异侧时，返回root
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
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