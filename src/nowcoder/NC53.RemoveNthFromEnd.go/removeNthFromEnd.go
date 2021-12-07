package NC53
import . "structs"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
  * 删除链表倒数第n个 双指针
  * @param head ListNode类 
  * @param n int整型 
  * @return ListNode类
*/
func removeNthFromEnd( head *ListNode ,  n int ) *ListNode {
    // write code here
	front,end := head,head
	//end 先行k步
	for n>1 && end!=nil{
		end = end.Next
		n--
	}
	//如果出现n>0 而end==nil  说明链表长度<n
	if n>0 && end==nil {
		return head
	}
	//长度=n
	if n==0 && end==nil {
		return head.Next
	}
	//同步
	pre := front
	for end!=nil {
		pre = front
		front = front.Next
		end = end.Next
	}
	//删除front
	pre.Next = front.Next
	front.Next = nil  // gc
	return head
}