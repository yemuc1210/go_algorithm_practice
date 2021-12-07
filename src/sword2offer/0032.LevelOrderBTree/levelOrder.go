package offer32
import "structs"
type TreeNode=structs.TreeNode
/*
树的层序遍历
提示：
节点总数 <= 1000
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //4ms 7.60%   2.6MB   77.33%
func levelOrder(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
	res :=[]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)   //入队
	for len(queue) > 0{
		front := queue[0]
		queue = queue[1:]    //出队  先进先出
		res = append(res, front.Val)
		//访问该节点的左右节点
		if front.Left != nil{
			queue = append(queue, front.Left)
		}
		if front.Right != nil{
			queue = append(queue, front.Right)
		}
	}
	return res
}