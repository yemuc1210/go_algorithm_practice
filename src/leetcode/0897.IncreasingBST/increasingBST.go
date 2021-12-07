package lt897

import (
	// "math"
	"structs"
)
type TreeNode = structs.TreeNode
/*
剑指 Offer II 052. 展平二叉搜索树
给你一棵二叉搜索树，请 按中序遍历 将其重新排列为一棵递增顺序搜索树，
使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。

algorithm:
	先对输入的二叉搜索树执行中序遍历，将结果保存到一个列表中；
	然后根据列表中的节点值，创建等价的只含有右节点的二叉搜索树，其过程等价于根据节点值创建一个链表。
*/
func increasingBST(root *TreeNode) *TreeNode {
	//首先进行中序遍历，得到结果
	vals := []int{}
	var inorder func(*TreeNode)
	inorder = func(tn *TreeNode) {
		if tn != nil {
			inorder(tn.Left)
			vals = append(vals, tn.Val)
			inorder(tn.Right)
		}
	}
	inorder(root)
	//构建BST
	dummyNode := &TreeNode{}

	curNode := dummyNode
	for _,val := range vals{
		curNode.Right = &TreeNode{Val: val}
		curNode = curNode.Right
	}
	return dummyNode.Right
}
