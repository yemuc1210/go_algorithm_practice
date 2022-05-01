package offerS78

import (
	// "sort"
	"go_practice/structs"
)

/* lt23
剑指 Offer II 078. 合并排序链表
给定一个链表数组，每个链表都已经按升序排列。
请将所有链表合并到一个升序链表中，返回合并后的链表。

归并
*/
type ListNode = structs.ListNode
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1{
		return nil
	}
	if length == 1{ //因为下面的合并函数需要两个子链
		return lists[0]  //一个链不要合并
	}

	num := length/2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoList(left,right)
}

func mergeTwoList(head1,head2 *ListNode)*ListNode{
	if head1 == nil{
		return head2
	}
	if head2 == nil {
		return head1
	}
	if head1.Val < head2.Val {
		head1.Next = mergeTwoList(head1.Next,head2) //递归
		return head1
	}
	head2.Next = mergeTwoList(head1,head2.Next)
	return head2
}