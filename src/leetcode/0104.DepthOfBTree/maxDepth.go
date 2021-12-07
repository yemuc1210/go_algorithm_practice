package lt104 
import (

	"structs"
)
type TreeNode=structs.TreeNode
/*
输入一棵二叉树的根节点，求该树的深度。
从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，
最长路径的长度为树的深度。

一般思路是层次遍历
*/
func maxDepth1(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if root.Left==nil && root.Right==nil{
		return 1
	}
	depth := 0
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0{
		length := len(queue)   //得到这一层的宽度
		for i:=0;i<length;i++{
			cur := queue[0]  //
			queue = queue[1:]  
			if cur.Left != nil{
				queue = append(queue, cur.Left)
			}
			if cur.Right !=nil{
				queue = append(queue, cur.Right)
			}
		}
		depth +=1
	}
	return depth
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}