package lt814
import "structs"
type TreeNode = structs.TreeNode
/*
剑指 Offer II 047. 二叉树剪枝
给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。
请剪除该二叉树中所有节点的值为 0 的子树。

节点 node 的子树为 node 本身，以及所有 node 的后代。

只能减除子树数全是0 所有不包含1
*/
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left==nil && root.Right==nil{
		return nil
	}
	return root
}