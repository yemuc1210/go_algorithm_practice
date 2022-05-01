package lt114
import (
	"go_practice/structs"
)
type TreeNode=structs.TreeNode
/* lt430 多链表展开
114. 二叉树展开为链表
给你二叉树的根结点 root ，请你将它展开为一个单链表：
	展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
	展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/
/*
非递归
*/
func flatten(root *TreeNode)  {
	list,cur := []int{}, &TreeNode{}
	preorder(root,&list)
	cur = root
	for i:=1;i<len(list);i++{
		cur.Left = nil
		cur.Right = &TreeNode{Val: list[i], Left: nil, Right: nil}
		cur = cur.Right
	}
	return 
}

func preorder(root *TreeNode, output *[]int)  {
	if root != nil {
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right,output)
	}
}