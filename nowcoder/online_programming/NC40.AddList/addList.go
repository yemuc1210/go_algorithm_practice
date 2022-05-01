package NC40
import . "go_practice/structs"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 两个链表生成相加链表 addTwoNumbers  
 * @param head1 ListNode类 
 * @param head2 ListNode类 
 * @return ListNode类
*/
func addInList( head1 *ListNode ,  head2 *ListNode ) *ListNode {
    // write code here 链表非逆序 可以转换成数组再做
	//得到两个数组 可以看作栈 计算时弹出栈顶（最后一个）
	s1,s2,carry := listToArr(head1),listToArr(head2),0
	var newHead *ListNode
	for len(*s1) > 0 || len(*s2) > 0 {
		carry += pop(s1) + pop(s2)
		newHead, carry = &ListNode{Val: carry % 10, Next: newHead}, carry/10
	}
	if carry != 0 {
		newHead = &ListNode{Val: carry, Next: newHead}
	}
	return newHead

}

func listToArr(head *ListNode) *[]int{
	stack := new([]int)  //返回类型为*【】int
	for ptr := head; ptr != nil; ptr = ptr.Next {
		*stack = append(*stack, ptr.Val)
	}
	return stack
}

func pop(stack *[]int)(pop int){
	if len(*stack) > 0 {
		pop = (*stack)[len(*stack)-1]
		*stack = (*stack)[0 : len(*stack)-1]
	}
	return
}