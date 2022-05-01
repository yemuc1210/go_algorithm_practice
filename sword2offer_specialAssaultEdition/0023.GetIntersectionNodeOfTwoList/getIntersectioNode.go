package offerS23

import "go_practice/structs"
type ListNode=structs.ListNode

/*
剑指 Offer II 023. 两个链表的第一个重合节点
给定两个单链表的头节点 headA 和 headB ，请找出并返回两个单链表相交的起始节点。
如果两个链表没有交点，返回 null 。

图示两个链表在节点 c1 开始相交：
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	//如果a b 有不同的长度，在第二个迭代后停止
	for a != b {
		//当短的链表到头后，指向另一个链表，此时两个指针距离=链表长度差
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		
	}
	return a
}