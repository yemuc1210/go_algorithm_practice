package lt95

import (
	"structs"
)
type TreeNode=structs.TreeNode
/*中 lt96不同的二叉搜索树  
95. 不同的二叉搜索树 II
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
*/
/*
有规律的吧  想不出
思路2：
	递归求解。外层循环遍历 1~n 所有结点，作为根结点，内层双层递归分别求出左子树和右子树。
*/
func generateTrees(n int) []*TreeNode {
	if n==0 {
		return []*TreeNode{}
	}
	var generateBST func(start,end int) []*TreeNode
	generateBST = func(start, end int) []*TreeNode {
		tree := []*TreeNode{}
		if start > end {
			tree = append(tree, nil)
			return tree
		}
		//i 作为根节点
		for i:=start;i<=end;i++{  //例子的数字是从1开始的n个数，所以范围闭区间[1,n]
			left := generateBST(start,i-1)
			right := generateBST(i+1,end)
			for _,l := range left {
				for _,r := range right {
					root := &TreeNode{Val:i, Left: l, Right: r}
					tree = append(tree, root)
				}
			}
		}
		return tree
	}
	return generateBST(1,n)
}
