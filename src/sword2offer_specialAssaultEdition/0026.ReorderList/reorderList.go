package offerS26

import "structs"
type ListNode=structs.ListNode

/*
剑指 Offer II 026. 重排链表
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

 L0 → L1 → … → Ln-1 → Ln 
请将其重新排列后变为：

L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …

不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
func reorderList(head *ListNode)  {
	//双指针、首 中间节点
	if head == nil || head.Next == nil {
		return 
	}

	// 寻找中间结点
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	// 反转链表后半部分  1->2->3->4->5->6 to 1->2->3->6->5->4
	preMiddle := p1
	preCurrent := p1.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

	// 重新拼接链表  1->2->3->6->5->4 to 1->6->2->5->3->4
	p1 = head
	p2 = preMiddle.Next
	for p1 != preMiddle {
		preMiddle.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = preMiddle.Next
	}
	return 
}