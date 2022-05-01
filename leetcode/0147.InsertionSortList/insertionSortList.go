package lt147

import (
	"go_practice/structs"
	"math"
)

type ListNode = structs.ListNode

/*
对链表进行插入排序
*/
func insertionSortList(head *ListNode) *ListNode {
	newHead := &ListNode{Val: math.MinInt32, Next: head}
	//使用head作为遍历指针
	head = head.Next
	newHead.Next.Next = nil //与原链表断开
	for head != nil {
		//将head插入newHead链表
		pre, cur := newHead, newHead.Next
		for cur != nil && cur.Val < head.Val {
			pre = cur
			cur = cur.Next
		}
		//p可能为nil
		if cur == nil { //说明当前head节点插在链表尾部
			pre.Next = head
			head = head.Next
			pre.Next.Next = nil //断开与其他链的连接
		} else { //当前head.Val <= p.Val  插在p的前面
			tmpHead := head.Next
			head.Next = cur
			pre.Next = head
			head = tmpHead
		}
	}
	return newHead.Next
}
