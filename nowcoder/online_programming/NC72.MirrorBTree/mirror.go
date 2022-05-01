package NC72
import  . "go_practice/structs"

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 二叉树的镜像  offer27
 * 将二叉树变为其镜像
 * @param pRoot TreeNode类 
 * @return TreeNode类
*/
func Mirror( root *TreeNode ) *TreeNode {
	if root == nil {
		return nil
	}
	left := Mirror(root.Left)
	right := Mirror(root.Right)
	root.Left = right
	root.Right = left
	return root
}

//使用栈
func mirror(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		//交换
		cur.Left,cur.Right = cur.Right,cur.Left
	}
	return root
}