package NC136

import (
	"fmt"
	"go_practice/structs"
)
type TreeNode = structs.TreeNode
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 求二叉树的右视图
 * 给定前序 中序 恢复二叉树 并打印右视图
 * @param xianxu int整型一维数组 先序遍历
 * @param zhongxu int整型一维数组 中序遍历
 * @return int整型一维数组
*/
func solve( preorder []int ,  inorder []int ) []int {
    // write code here
	root := PreIn2Tree(preorder,inorder)
	return rightSideView(root)
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
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

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
//BFS 记录每层最后一个
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue)>0 {
		size := len(queue)    //每层的宽度
		for i:=0;i<size;i++ {
			node := queue[0]   //出队
			queue = queue[1:]
			if node.Left != nil{
				queue = append(queue, node.Left)
			}
			if node.Right != nil{
				queue = append(queue, node.Right)
			}
			if i== size-1 {
				res = append(res, node.Val)
			}
		}
	}
	return res
}