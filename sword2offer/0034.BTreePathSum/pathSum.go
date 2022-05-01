package offer34

import "go_practice/structs"
type TreeNode = structs.TreeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。
 		从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
 */

//dfs完事
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res:=[][]int{}
	
	findPath(root,target,&res,[]int(nil))
	return res
}

func findPath(root *TreeNode,target int, res *[][]int,stack []int){
	if root==nil{
		return 
	}
	target -= root.Val
	stack = append(stack, root.Val)
	if target==0 && root.Left==nil && root.Right==nil{
		*res = append(*res, append([]int{},stack...))
		stack = stack[:len(stack)-1]
	}
	findPath(root.Left,target,res,stack)
	findPath(root.Right,target,res,stack)
}