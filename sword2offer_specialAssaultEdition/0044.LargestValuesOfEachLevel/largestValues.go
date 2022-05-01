package offerS44

import "go_practice/structs"
type TreeNode = structs.TreeNode

/*
剑指 Offer II 044. 二叉树每层的最大值
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
*/
func largestValues(root *TreeNode) []int {
	//处理空
	if root == nil {
		return []int{}
	}

	//层次遍历
	res := []int{}
	q,curMax,curLevelNum := []*TreeNode{root},0,0

	for len(q) > 0 {
		//统计
		curLevelNum = len(q)   //当前队列长度，也就是某一层的宽度
		curMax = q[0].Val
		for i:=0;i<curLevelNum;i++{
			//当前遍历层出队
			curNode := q[0]
			q = q[1:]

			if curNode.Val > curMax {
				curMax = curNode.Val
			}

			//左右节点入队
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
		}
		//统计完当前层，最大值入结果
		res = append(res, curMax)
	}

	return res
}