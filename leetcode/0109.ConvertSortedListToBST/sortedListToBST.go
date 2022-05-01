package lt109
//lt108

import (
	"go_practice/structs"
)
type ListNode=structs.ListNode
type TreeNode = structs.TreeNode
/* 中 
109. 有序链表转换二叉搜索树
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
*/
/*
108有序数组转BST
109有序链表
有序数组的思路是中间节点做根，左右两边递归构建子树
有序链表取中间节点，使用快慢指针，快指针每次走两步
思路1：有序链表转有序数组，然后套用108方法
思路2：快慢指针取中点，递归  更优
*/
func sortedListToBST(head *ListNode) *TreeNode {
	nums := listToArr(head)
	return arrToBST(nums)
}
func arrToBST(nums []int) *TreeNode{
	if len(nums) == 0{
		return nil
	}
	return &TreeNode{
		Val: nums[len(nums)/2],
		Left: arrToBST(nums[:len(nums)/2]),
		Right: arrToBST(nums[len(nums)/2+1:]),
	}
}
func listToArr(head *ListNode) []int{
	if head == nil {
		return []int{}
	}
	res := []int{}
	for head !=nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func sortedListToBST1(head *ListNode) *TreeNode{
	if head == nil {
		return nil
	}else if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	//快慢指针取中点  此时链表长度为2
	pre := head
	p := pre.Next  //所以p不为空  慢指针
	q := p.Next  //快指针
	//找到中点  p
	for q!=nil && q.Next != nil {
		pre = pre.Next
		p = pre.Next
		q = q.Next.Next
	}
	//中点分左右
	pre.Next = nil   //左分
	root := &TreeNode{Val: p.Val}
	root.Left = sortedListToBST1(head)
	root.Right = sortedListToBST1(p.Next)
	return root
}