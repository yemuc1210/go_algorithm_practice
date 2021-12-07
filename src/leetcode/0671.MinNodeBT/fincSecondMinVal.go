package lt671

import "structs"

type TreeNode = structs.TreeNode

/*
每个节点的子节点数量为0、2，如果有两个子节点，则该节点的值为两节点中较小的
输出所有节点中第二小的，否则-1

也就是求所有叶子节点中第二小的

根节点值是最小的，通过遍历找出严格大于根节点值的最小值
遍历  dfs
*/
func findSecondMinimumValue(root *TreeNode) int{
	res := -1
	rootVal := root.Val
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || res != -1 && node.Val >= res{
			return 
		}
		if node.Val > rootVal {
			res = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}