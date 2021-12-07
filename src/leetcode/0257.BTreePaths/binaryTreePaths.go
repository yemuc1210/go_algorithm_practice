package lt257

import (
	"strconv"
	"structs"
)
type TreeNode = structs.TreeNode
/*
257. 二叉树的所有路径
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。
*/
/*
dfs
*/
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	dfs(root, &res, "")
	return res
}
func dfs(root *TreeNode, res *[]string, path string) {
	path += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return 
	}
	if root.Left!=nil {
		dfs(root.Left, res, path+"->")
	}
	if root.Right!=nil{
		dfs(root.Right, res, path+"->")
	}
	
}	