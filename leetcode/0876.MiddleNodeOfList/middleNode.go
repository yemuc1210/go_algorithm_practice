package lt876

import "go_practice/structs"

type ListNode = structs.ListNode

/*
给定一个头结点为 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。


首尾指针 碰撞   不行，无法向前
快慢指针  2 1
*/
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next==nil{
		return head
	}

	fast,slow := head,head
	//快2  慢1
	for fast != nil && fast.Next!=nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}