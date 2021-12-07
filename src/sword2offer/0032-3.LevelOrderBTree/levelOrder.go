package offer32

import (
	// "sort"
	"structs"
)
type TreeNode=structs.TreeNode
/*
请实现一个函数按照之字形顺序打印二叉树，
	即第一行按照从左到右的顺序打印，
	第二层按照从右到左的顺序打印，
	第三行再按照从左到右的顺序打印，其他行以此类推。
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
 //0ms 100%   2.8MB   50.30%
//我特么直接数组排序
func levelOrder(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	res :=[][]int{}
	curNum,nextNum,tmp :=1,0,[]int{}

	flag := 0
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
			if flag%2==0 {
				res = append(res, tmp)
				flag++
			}else{
				reverse(tmp)
				res = append(res, tmp)
				flag++
			}
			curNum=nextNum
			nextNum=0
			tmp=[]int{}
		}
		
	}
	return res
}
func reverse(arr []int){
	n:=len(arr)
	for i:=0;i<n/2;i++{
		arr[i],arr[n-i-1] = arr[n-i-1],arr[i]
	}
}