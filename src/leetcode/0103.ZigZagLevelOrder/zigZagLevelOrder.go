/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 package lt103
 import "structs"
type TreeNode = structs.TreeNode
//剑指Offer 32-3
 func zigzagLevelOrder(root *TreeNode) [][]int {
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