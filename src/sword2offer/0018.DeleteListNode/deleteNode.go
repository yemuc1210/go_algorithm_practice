package offer18

import "structs"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ 
/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。
题目保证链表中节点值各不相等
*/
type ListNode=structs.ListNode
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil{
		return nil
	}
	//处理头结点就是要删除的情况
	if head.Val==val{
		p:=head
		head = head.Next
		p.Next=nil
		return head
	}
	p,q := head,head
	for p!=nil&& p.Val!=val{
		q = p
		p = p.Next
	}

		q.Next = p.Next
		p.Next = nil

	return head
}