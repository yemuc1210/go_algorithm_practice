package NC51
import . "go_practice/structs"

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 合并k个已排序的链表  两两合并即可 lt23
 * @param lists ListNode类一维数组 
 * @return ListNode类
*/
func mergeKLists( lists []*ListNode ) *ListNode {
    if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1{
		return lists[0]
	}
	num := len(lists) / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists(left, right)
}
func mergeTwoLists(l1,l2 *ListNode) *ListNode{
	if l1==nil {
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1,l2.Next)
	return l2
}