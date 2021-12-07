package offer27
import "structs"

type TreeNode = structs.TreeNode

/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
*/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func mirrorTree1(root *TreeNode) *TreeNode{
	//辅助栈
	if root == nil {
		return nil
	}

	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		node.Left,node.Right = node.Right,node.Left
	}
	return root
}