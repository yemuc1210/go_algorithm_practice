package lt106

import (
	"fmt"
	"structs"
)
type TreeNode = structs.TreeNode
/*
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/
/*
后序遍历最后一个元素是root  到中序遍历找到root，恰好分割左右子树
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) != len(inorder) {
		panic("inPost2Tree 中两个切片的长度不相等")
	}

	if len(inorder) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: postorder[len(postorder)-1],   
	}  //首先构建root

	if len(inorder) == 1 {
		return res
	}

	idx := indexOf(res.Val, inorder)

	res.Left = buildTree(inorder[:idx], postorder[:idx]) //递归构建左右子树
	res.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])

	return res
}

//查找元素val在数组nums中的下标
func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	panic(fmt.Sprintf("%d 不存在于 %v 中", val, nums))
}