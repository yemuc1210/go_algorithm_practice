package lt102

import (
	"go_practice/structs"
)
type TreeNode = structs.TreeNode


func levelOrder(root *TreeNode) [][]int{
	if root == nil{
		return [][]int{}
	}

	queue := []*TreeNode{}
	width := 1 // 记录每一层的宽度
	//根入队
	queue = append(queue, root)
	res :=[][]int{}
	for len(queue)>0{
		width = len(queue)
		tmpArr := []int{}
		for i:=0;i<width;i++{
			
			cur := queue[0]
			queue = queue[1:]
			tmpArr = append(tmpArr, cur.Val)
			if cur.Left!=nil{
				queue = append(queue, cur.Left)
			}
			if cur.Right!=nil{
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, tmpArr)
	}
	return res

}