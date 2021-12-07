package lt513

import "structs"
type TreeNode = structs.TreeNode

/*
513. 找树左下角的值
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

假设二叉树中至少有一个节点。
*/
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	//BFS即可
	q,curLeft,curLevelNum := []*TreeNode{root},root.Val,0

	for len(q) > 0 {
		curLeft = q[0].Val    //每一层第一个就是当前层最左节点
		curLevelNum = len(q)

		for i:=0;i<curLevelNum;i++{
			//处理当前层
			curNode := q[0]
			q = q[1:]

			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
		}
	}
	return curLeft
}