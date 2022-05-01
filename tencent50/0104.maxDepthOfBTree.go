package tencent50
import "go_practice/structs"
type TreeNode = structs.TreeNode
/*简单
104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。
*/
/*
递归：左右子树深度的较大值

非递归：bfs
*/
//使用bfs求解
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevelWidth := len(queue)
		//遍历完当前层
		for i:=0;i<curLevelWidth;i++{
			curRoot := queue[0]
			queue = queue[1:]  //出队 队头出
			//左右节点入队
			if curRoot.Left != nil {
				queue = append(queue, curRoot.Left)
			}
			if curRoot.Right != nil {
				queue = append(queue, curRoot.Right)
			}
		}
		depth ++
	}
	return depth
}