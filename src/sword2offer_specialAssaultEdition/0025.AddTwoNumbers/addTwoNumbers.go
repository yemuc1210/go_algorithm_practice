package offerS25
import "structs"
type ListNode=structs.ListNode
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2, carry := pushList(l1), pushList(l2), 0
	var newHead *ListNode
	for len(*s1) > 0 || len(*s2) > 0 {
		carry += popStack(s1) + popStack(s2)
		newHead, carry = &ListNode{Val: carry % 10, Next: newHead}, carry/10
	}
	if carry != 0 {
		newHead = &ListNode{Val: carry, Next: newHead}
	}
	return newHead
}

func pushList(head *ListNode) *[]int {
	stack := new([]int)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		*stack = append(*stack, ptr.Val)
	}
	return stack
}

func popStack(stack *[]int) (pop int) {
	if len(*stack) > 0 {
		pop = (*stack)[len(*stack)-1]
		*stack = (*stack)[0 : len(*stack)-1]
	}
	return
}

