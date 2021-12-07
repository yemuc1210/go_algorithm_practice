package lt203
import "structs"
type ListNode = structs.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
给你一个链表的头节点 head 和一个整数 val ，
请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0,Next: head}
	pre,cur := newHead,head
	for cur!=nil{
		if cur.Val==val{
			pre.Next = cur.Next
		}else{
			pre=cur
		}
		cur = cur.Next
	}
	return newHead.Next
}