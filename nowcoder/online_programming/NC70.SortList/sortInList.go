package NC70

import (
	"math"
	. "go_practice/structs"
)

/**
 * 单链表排序  升序排
 * @param head ListNode类 the head node
 * @return ListNode类
 */
func sortInList( head *ListNode ) *ListNode {
    // 升序排列单链表  插入排序的思路
	if head == nil || head.Next == nil{
		return head
	}
	newHead := &ListNode{Val: math.MinInt64, Next: head}
	p := head.Next  //last处理过的链表最后一个节点
	newHead.Next.Next = nil // 断开连接
	for p!=nil {
		//未排序链表第一个节点  p始终指向第一个
		//为p寻找插入位置
		curNode := newHead
		for curNode.Next!=nil && curNode.Next.Val < p.Val{
			curNode = curNode.Next  //升序
		}
		//可能curNode.Next == nil 那么在末尾插入
		//跳出时  curNode后面插入
		tmpNode := p.Next
		p.Next = curNode.Next
		curNode.Next = p
		p = tmpNode
	}
	return newHead.Next
}