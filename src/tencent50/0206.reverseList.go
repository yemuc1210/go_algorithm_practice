package tencent50
/*简单
206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/

/*
头插法
*/
func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	curNode := head

	for curNode != nil {
		//头插
		tmp := curNode.Next
		curNode.Next = newHead.Next
		newHead.Next = curNode
		curNode = tmp

	}
	
	return newHead.Next
}