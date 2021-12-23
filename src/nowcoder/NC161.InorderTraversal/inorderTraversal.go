package NC161
import . "structs"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 中序遍历
 * 
 * @param root TreeNode类 
 * @return int整型一维数组
*/
func inorderTraversal( root *TreeNode ) []int {
    // write code here
	stack := []*TreeNode{}
	res := []int{}

	for root!=nil || len(stack) >0 {
		for root!=nil {
			//根
			stack = append(stack, root)
			//左
			root = root.Left
		}  //存到栈就表示 待遍历节点
		//遍历左节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)

		//右
		root = root.Right
	}
	return res
}

func inorder(root *TreeNode) [] int{
	stack := []*TreeNode{}
	res := []int{}

	for root!=nil || len(stack) >0 {
		for root!=nil {
			//根
			stack = append(stack, root)
			//左
			root = root.Left
		}  //存到栈就表示 待遍历节点
		//遍历左节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)

		//右
		root = root.Right
	}
	return res

}