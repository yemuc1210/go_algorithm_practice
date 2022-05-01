package lt86
import "go_practice/structs"
type ListNode = structs.ListNode

/*
86. 分隔链表
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，
使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。
*/
/*
保持初始相对位置：尾插
找到小于x的节点 尾插到小于x链表后边
*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	lessList := &ListNode{}
	tail := lessList
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	cur := dummyHead.Next
	for cur!=nil {
		if cur.Val < x {
			//断开cur 前面的
			prev.Next = cur.Next
			//现在可以插cur到lessList
			tail.Next = cur
			cur.Next = nil
			tail = cur	

			cur = prev.Next
			continue
		}
		prev = cur
		cur = cur.Next
	}
	tail.Next = dummyHead.Next
	return lessList.Next
}