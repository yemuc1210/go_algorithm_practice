package NC102
import . "go_practice/structs"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 找到 o1 和 o2 的最近公共祖先节点  lt235 lt236
 * 若 root 为o1 o2的最近公共祖先，存在以下集中情况
	（1）o1 o2 在异侧
	（2）否则 root必定是其中一个
		1）root.Val = o1   o2在左子树或右子树
		2）root.Val = o2  o1在左子树或右子树
 * @param root TreeNode类 
 * @param o1 int整型 
 * @param o2 int整型 
 * @return int整型
*/
func lowestCommonAncestor( root *TreeNode ,  o1 int ,  o2 int ) int {
    // write code here  普通二叉树 非二叉搜索树
	if root==nil{
		return 0
	}
	if root.Val == o1 || root.Val == o2{
		return root.Val  //题目的数据范围是【1，1000】
	}
	left := lowestCommonAncestor(root.Left, o1, o2)
	right := lowestCommonAncestor(root.Right, o1,o2)
	if left>0 {
		if right >0 {
			return root.Val
		}
		return left
	}

	return right
}