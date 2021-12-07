package offer28

import "structs"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
请实现一个函数，
用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
限制：
0 <= 节点个数 <= 1000
*/

//0ms  100%   2.9MB  16.19%
type TreeNode=structs.TreeNode
func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	return isSameTree(invertTree(root.Left), root.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	//遍历改编
	if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}