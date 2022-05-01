package NC81
import . "go_practice/structs"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 二叉搜索树的第k个节点   offer54
 * 第k小   中序遍历
 * @param proot TreeNode类 
 * @param k int整型 
 * @return int整型
*/
func KthNode( proot *TreeNode ,  k int ) int {
    // write code here
	order := inorderTraversal(proot)
	//条件0<=k<=n
	n := len(order)  
	if len(order) == 0{
		return -1
	}
	if k>n {
		return -1
	}
	if k==0 {
		return -1
	}
	return order[k-1]
}

//迭代法 用到栈
func inorderTraversal(root *TreeNode)[]int{
	res := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)  //根节点入栈
			root = root.Left  //迭代访问左节点
		}
		root = stack[len(stack)-1]   //最后一个就是最左 根 然后访问右节点
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
	
		root = root.Right   //访问右节点
	}
	return res
}