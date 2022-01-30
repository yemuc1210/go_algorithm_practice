package NC3
import . "structs"
/*
链表中环得入口节点  offerS22  lt142
当找到环时，返回此时的相遇节点
分析一下判断环的原理。fast 指针一次都 2 步，slow 指针一次走 1 步。
令链表 head 到环的一个点需要 x1 步，从环的第一个点到相遇点需要 x2 步，
从环中相遇点回到环的第一个点需要 x3 步。
那么环的总长度是 x2 + x3 步。
fast 和 slow 会相遇，说明他们走的时间是相同的，可以知道他们走的路程有以下的关系：
fast 的 t = (x1 + x2 + x3 + x2) / 2
slow 的 t = (x1 + x2) / 1
x1 + x2 + x3 + x2 = 2 * (x1 + x2)
	所以 x1 = x3
然后fast从头开始
*/
func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	if pHead == nil || pHead.Next  == nil{
		return nil
	}
	isCycle,slow := hasCycle(pHead)
	if !isCycle {
		return nil
	}
	fast := pHead
	//同步
	for fast!=slow{
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

//使用快慢指针寻找环
func hasCycle(head *ListNode)(bool, *ListNode) {
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