package main

import "go_practice/structs"

type TreeNode = structs.TreeNode

func isUnivalTree(root *TreeNode) bool {
	// 单值二叉树
	// 每个节点都是相同的值
	if root == nil {
		return true
	}
	// bfs
	q := []*TreeNode{root}
	uniqueVal := root.Val
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.Val != uniqueVal {
			return false
		}
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right !=nil {
			q = append(q, cur.Right)
		}
	}
	return true
}