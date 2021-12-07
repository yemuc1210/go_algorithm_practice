package NC45
import . "structs"

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 实现二叉树的三种遍历 要求时空复杂度O(n)
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
*/
func threeOrders( root *TreeNode ) [][]int {
    // write code here
	return [][]int{
		postOrder(root),inOrder(root),postOrder(root),
	}
}
func preOrder(root *TreeNode) []int {
	res := []int{}	
	preorder(root,&res)
	return res
}

func preorder(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right, output)
	}
}


func inOrder(root *TreeNode) []int {
	//迭代法 lt94
	stack,res := []*TreeNode{}, []int{}
	for root!=nil || len(stack) >0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func postOrder(root *TreeNode)[]int{
	//后序迭代的递归有点难
	res := []int{}
	postorder(root, &res)
	return res
}
func postorder(root *TreeNode, output *[]int) {
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		*output = append(*output, root.Val)
	}
}