package  interview04

import "go_practice/structs"
type TreeNode = structs.TreeNode

/*
面试题 04.02. 最小高度树
给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。


BST  升序  中间的作为中间节点  二分
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}

	if len(nums)==1{
		return &TreeNode{Val: nums[0]}
	}

	mid := len(nums)/2

	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}