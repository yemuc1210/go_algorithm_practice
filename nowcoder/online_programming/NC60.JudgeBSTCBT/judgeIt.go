package NC60

import (
	"math"
	. "go_practice/structs"
)

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 判断二叉树是否是二叉搜索树和完全二叉树
 * @param root TreeNode类 the root
 * @return bool布尔型一维数组
 */
func judgeIt( root *TreeNode ) []bool {
    // write code here
	var res = make([]bool, 2)
	res[0] = isBST(root,math.MinInt32,math.MaxInt32)
	res[1] = isCBT(root)
	return res
}
//判断是否是BST 
func isBST(root *TreeNode, left,right int) bool {
	if root == nil {
		return true
	}  //空子树认为符合二叉搜索树
	if root.Val <= left || root.Val >= right {
		return false
	}
	return isBST(root.Left,left,root.Val) && isBST(root.Right, root.Val, right)
}
//判断是否是完全二叉树 BFS
func isCBT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		
		//如果遇到nil 那么之后得元素必须都是nil
		if curNode == nil {
			//查看队列元素是否都是nil
			for i:=0;i<len(queue);i++{
				if queue[i] != nil {
					return false
				}
			}
			return true
		}
		queue = append(queue, curNode.Left)
		queue = append(queue, curNode.Right)
	}
	return true
}	