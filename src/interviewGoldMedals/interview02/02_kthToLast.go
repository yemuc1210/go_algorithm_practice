package interview02
/*
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

说明：k有效   双100%
*/
func kthToLast(head *ListNode, k int) int {
	//特殊情况判断
	//空
	if head == nil {
		return -1
	}
	if head != nil && head.Next == nil{
		if k == 1{
			return head.Val
		}
		return -1
	}
	//双指针
	fast,slow := head,head

	//fast先行k步
	for steps:=0; steps < k; steps++{
		fast = fast.Next
		if fast == nil{
			if steps == k-1 { //k正好是链表长度
				return slow.Val
			}else{
				return -1
			}
		}
	}
	//现在fast先k步
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val

}