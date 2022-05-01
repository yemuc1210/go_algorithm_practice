package NC25
import . "go_practice/structs"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
  * 删除有序链表中的重复元素
  * 删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次
  * 与24的区别在于，重复元素不删除
  * @param head ListNode类 
  * @return ListNode类
*/
func deleteDuplicates( head *ListNode ) *ListNode {
    if head == nil {
		return head
	}
	dummyHead := &ListNode{Val :-101, Next: head}
	//思路：pre  next为下一个不同的元素 然后删去中间的
	pre,next := dummyHead,dummyHead.Next
	for pre != nil { //pre也就是当前待处理元素
		next = pre.Next
		for next!=nil &&next.Val==pre.Val {
			next =  next.Next
		}
		//此时next.Val != pre.Val
		//next可能为nil
		if next == nil {
			//此时pre应该是结果链表的最后一个元素 其后元素都是重复的
			pre.Next = nil
		}
		//否则   删除中间重复的
		pre.Next = next
		pre = next
	}
	return dummyHead.Next
}