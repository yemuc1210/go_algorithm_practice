package NC69

import . "go_practice/structs"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 链表倒数最后k节点  offer22
 * 
 * @param pHead ListNode类 
 * @param k int整型 
 * @return ListNode类
*/
func FindKthToTail( head *ListNode ,  k int ) *ListNode {
    	//双指针
	former,latter:=head,head
	for i:=0;i<k;i++{
		if former==nil{
			return nil
		}
		former=former.Next
	}
	for former!=nil{
		former=former.Next
		latter=latter.Next
	}
	return latter
}