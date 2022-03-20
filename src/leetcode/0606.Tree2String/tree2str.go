package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 采用前序遍历的方式，将二叉树转换成一个由括号和整数组成的字符串
// a(b(c))  表示a具有b子节点，c孙节点、
func tree2str(root *TreeNode) {
	// 每访问一个阶段，输出(node.Val，return 时+)
	// 使用前序遍历的迭代
	stack := []*TreeNode{}
	stack = append(stack, root)
	var res []interface{}
	res = append(res, "(")
	res = append(res, root.Val)
	for len(stack) > 0 {
		// 前序遍历的顺序是 中左右
		// 访问
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, ")")
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
			res = append(res, "(")
			res = append(res, curNode.Right.Val)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
			res = append(res, "(")
			res = append(res, curNode.Left.Val)
		}
	}
	fmt.Println(res...)
	// return string(res)
}
func tree2str1(root *TreeNode) string {
	switch {
	case root == nil:
		return ""
	case root.Left == nil && root.Right == nil:
		return strconv.Itoa(root.Val)
	case root.Right == nil:
		return fmt.Sprintf("%d(%s)", root.Val, tree2str1(root.Left))
	default:
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str1(root.Left), tree2str1(root.Right))
	}
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}

	str := tree2str1(root)
	fmt.Println(str)
}
