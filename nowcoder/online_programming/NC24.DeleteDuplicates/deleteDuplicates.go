package NC24
import . "go_practice/structs"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
  * 删除有序链表中重复元素
  * @param head ListNode类 
  * @return ListNode类
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{Val: 0, Next: head}
	head = newHead

	lastVal := 0  //有序非重复 最后一个
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val { //后两个是重复
			lastVal = head.Next.Val
			for head.Next != nil && lastVal == head.Next.Val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return newHead.Next
}
