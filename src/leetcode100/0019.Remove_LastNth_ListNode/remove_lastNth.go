package leetcode 

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "structs"
type ListNode = structs.ListNode
/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p,q := head,head
	//首先将q先移动n个
	step:=0
	for i:=0;i<n;i++{
		//可能n比链表长度还大
		if q.Next==nil && step!=0&&step<n-1{
			return head
		}
		q = q.Next
		step++
	}
	if q == nil{ 
		head = head.Next
		return head
	}
	//那么此时q比p先n个距离，只要q到达最后一个，则p就是倒数第n个
	for q.Next != nil{
		p = p.Next
		
		q = q.Next
	}
	//此时，q应该指向最后一个
	p.Next = p.Next.Next   //移除位置是p的下一个
	return head
}