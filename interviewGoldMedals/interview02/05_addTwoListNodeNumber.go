package interview02
/*
面试题 02.05. 链表求和

给定两个用链表表示的整数，每个节点包含一个数位。

这些数位是反向存放的，也就是个位排在链表首部。

编写函数对这两个整数求和，并用链表形式返回结果。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	result := &ListNode{}
	dummy := result
	tmp := 0
	for l1 != nil || l2 != nil {
		result_v := tmp  //首先加上进位
		if l1 != nil {
			result_v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			result_v += l2.Val
			l2 = l2.Next
		}
		result.Next = &ListNode{Val: result_v % 10}
		tmp = result_v / 10  //进位

		result = result.Next
	}
	if tmp != 0 {
		result.Next = &ListNode{Val: tmp}
	}
	return dummy.Next
}