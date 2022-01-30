package NC23

import . "structs"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
  * 
  * @param head ListNode类 
  * @param x int整型 
  * @return ListNode类
*/
func partition( head *ListNode ,  x int ) *ListNode {
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