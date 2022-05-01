package lt111
import (
	"go_practice/structs"
)
type TreeNode=structs.TreeNode

/*简单   lt102层序遍历 lt104最大深度  剑指offer55
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。
*/
/*
思路还是用层序遍历
当遍历某一层时遇到左右子树都为空，则该层就是最小深度 返回深度即可
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}
	for len(queue) > 0{
		//遍历需要一层一层的遍历，所以需要知道每一层的宽度
		levelWidth := len(queue)
		//depth ++  则depth表示需要遍历的当前层深度
		depth += 1
		for i:=0;i<levelWidth;i++{
			curNode := queue[0]
			queue = queue[1:]
			//左右节点入队
			if curNode.Left != nil || curNode.Right != nil {
				//则该节点可以入队，不是叶子节点
				if curNode.Left != nil {
					queue = append(queue, curNode.Left)
				}
				if curNode.Right != nil {
					queue = append(queue, curNode.Right)
				}
			} else if curNode.Left==nil && curNode.Right==nil{ //根据短路 此分支进入 左右子树皆为nil
				return depth
			}
		}
		
	}
	return depth
}