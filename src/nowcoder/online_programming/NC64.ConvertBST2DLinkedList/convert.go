package NC64
//剑指offer 36
//将BST转换成排序的煦暖双向链表

type TreeNode struct {
	Val int
	Left *TreeNode  //双链表前指
	Right *TreeNode  //双链表后面的节点
}
/*
BST的中序遍历是顺序的
*/

var pre *TreeNode //必须在全局变量上才可以实现

/**
 * 二叉搜索树转换成排序的双向链表
 * @param pRootOfTree TreeNode类 
 * @return TreeNode类
*/


func Convert(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	p := root
	for p.Left!=nil {
		p = p.Left  //双向链表第一个
	}
	inorder(root)
	return p
}
func inorder(root *TreeNode) {
	if root==nil{
		return 
	}
	inorder(root.Left)

	root.Left = pre
	if pre != nil {
		pre.Right = root
	}
	pre = root // 迭代更新

	inorder(pre.Right)
}