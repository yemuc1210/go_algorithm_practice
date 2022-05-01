package lt107
import "go_practice/structs"
type TreeNode = structs.TreeNode
/*中 lt102 
107. 二叉树的层序遍历 II
给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
*/
/*
逆序遍历  用到栈  每次将一层的压入栈  然后输出栈结果即可
*/
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue,stack := []*TreeNode{root},[][]int{}
	for len(queue) > 0{
		levelCount := len(queue)  //当前层个数
		curLevel := []int{}
		for i:=0;i<levelCount;i++{
			curNode := queue[0]
			queue = queue[1:]  //队 先进先出
			//当前层入栈
			curLevel = append(curLevel, curNode.Val)
			//下一层入队
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		//当前层入栈
		stack = append(stack, append([]int(nil), curLevel...))
	}
	//栈 逆序
	n := len(stack)
	for i := 0; i<n/2 ;i++{
		stack[i],stack[n-i-1] = stack[n-i-1],stack[i]
	}
	return stack
}