package offer07

import (
	"fmt"
	"structs"
)
type TreeNode=structs.TreeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。

假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	return PreIn2Tree(preorder,inorder)
}
// PreIn2Tree 把 preorder 和 inorder 切片转换成 二叉树
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}    //先序第一个都是根节点

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)   //找到根节点在中序序列中的位置，划分左右字数

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])  //左子树
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])  //右子树

	return res
}
func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}