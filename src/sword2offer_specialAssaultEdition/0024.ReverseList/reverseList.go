package offerS24
import "structs"
type ListNode=structs.ListNode
/*
反转链表
*/
func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{Val: 0}
	//头插法
	if head == nil || head.Next==nil{
		return head
	}
	//否则head起码有两个元素 
	newHead.Next = head
	p,q := head.Next,head.Next  //p始终指向第一个
	head.Next=nil
	//把q插到p之前
	for p!=nil{
		q = p.Next
		p.Next=newHead.Next
		newHead.Next=p
		p=q
	}

	return newHead.Next
}