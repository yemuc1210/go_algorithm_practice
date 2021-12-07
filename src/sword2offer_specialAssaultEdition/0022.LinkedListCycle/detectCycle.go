package offerS22
import "structs"
type ListNode=structs.ListNode
/*
剑指 Offer II 022. 链表中环的入口节点
给定一个链表，返回链表开始入环的第一个节点。
 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。
 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 
如果 pos 是 -1，则在该链表中没有环。
注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

寻找环  快慢指针
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	isCycle, slow := hasCycle142(head)
	if !isCycle {
		return nil
	}
	fast := head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func hasCycle142(head *ListNode) (bool, *ListNode) {
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true, slow
		}
	}
	return false, nil
}
