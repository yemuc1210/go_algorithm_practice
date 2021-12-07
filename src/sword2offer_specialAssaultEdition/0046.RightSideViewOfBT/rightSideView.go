package offerS46

import "structs"
type TreeNode = structs.TreeNode

/*
剑指 Offer II 046. 二叉树的右侧视图
给定一个二叉树的 根节点 root，想象自己站在它的右侧，
按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

BFS 最右侧值
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	//BFS即可
	q,rightSide,curLevelNum := []*TreeNode{root},[]int{},0

	for len(q) > 0 {
		curLevelNum = len(q)

		for i:=0;i<curLevelNum;i++{
			//处理当前层

			curNode := q[0]
			q = q[1:]
			if i == curLevelNum-1 {
				//最右节点入
				rightSide = append(rightSide, curNode.Val)
			}
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
		}
	}
	return rightSide
}