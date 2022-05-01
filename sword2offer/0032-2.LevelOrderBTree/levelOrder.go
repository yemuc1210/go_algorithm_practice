package offer32
import "go_practice/structs"
type TreeNode=structs.TreeNode
/*
树的层序遍历
每一层输出到一行
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
 //0ms 100%   2.8MB   31,8%
func levelOrder(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	res :=[][]int{}
	curNum,nextNum,tmp :=1,0,[]int{}

	queue := []*TreeNode{}
	queue = append(queue, root)   //入队
	for len(queue) != 0{
		if curNum>0{
			front := queue[0]
			//访问该节点的左右节点
			if front.Left != nil{
				queue = append(queue, front.Left)
				nextNum++
			}
			if front.Right != nil{
				queue = append(queue, front.Right)
				nextNum++
			}
			curNum--
			tmp = append(tmp, front.Val)
			queue = queue[1:]    //出队  先进先出
		}
		if curNum==0{
			res = append(res, tmp)
			curNum=nextNum
			nextNum=0
			tmp=[]int{}
		}
		
	}
	return res
}