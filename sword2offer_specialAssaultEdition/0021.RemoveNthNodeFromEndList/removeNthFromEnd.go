package offerS21

import "go_practice/structs"
type ListNode=structs.ListNode

/*
剑指 Offer II 021. 删除链表的倒数第 n 个结点
给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	//可能存在n>length(list)的情况
	//首先得到链表长度
	fast,slow := head,head
	//fast  先行 n步
	
	for fast.Next != nil && n>0{
		fast = fast.Next
		n --
	}
	if n!= 0{  //则跳出条件为fast.Next == nil
		//说明n > length(list)  删除第一个节点
		return head.Next
	}
	// 下面fast slow同时前进  fast快n步
	for fast.Next != nil{
		fast = fast.Next
		slow = slow.Next
	}

	//fast.Next == nil   fast指向最后一个
	//移除slow的下一个节点
	slow.Next = slow.Next.Next
	return head

}