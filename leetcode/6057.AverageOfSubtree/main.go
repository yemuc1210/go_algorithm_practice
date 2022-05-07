package main

import (
	"fmt"
	"go_practice/structs"
)

type TreeNode = structs.TreeNode

// 节点值等于子树中值得平均值
// 获得子树值：后序遍历

// type TreeNode struct{
// 	Val int
// 	Left,Right *TreeNode
// }

func averageOfSubtree(root *TreeNode) int {
	var res int
	var postTraver func(node *TreeNode) (int, int)
	postTraver = func(node *TreeNode) (int, int) {
		if node != nil {
			var sum int
			var cnt int
			sum1, cnt1 := postTraver(node.Left)
			sum2, cnt2 := postTraver(node.Right)
			// 访问当前节点
			curNum := node.Val
			// if cnt2+cnt1 != 0 {
			// 	if curNum == (sum1+sum2)/(cnt1+cnt2) {
			// 		// 得到一组解
			// 		res++
			// 	}
			// }
			cnt = cnt1 + cnt2 + 1
			sum = sum1 + sum2 + curNum
			if curNum == sum/cnt {
				res ++
			}
			return sum, cnt
		}
		return 0, 0
	}
	postTraver(root)
	return res
}

func main() {
	nums := []int{4, 8, 5, 0, 1, structs.NULL, 6}
	root := structs.Ints2TreeNode(nums)
	res := averageOfSubtree(root)
	fmt.Println(res)
}
