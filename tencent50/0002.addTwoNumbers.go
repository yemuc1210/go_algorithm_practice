package tencent50
import 	"go_practice/structs"

type ListNode = structs.ListNode
/* 中
2. 两数相加
给你两个 非空 的链表，表示两个非负的整数。
它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/
/*
逆序存放，所以第一个是最低位，最后一个是最高位
需要考虑长度不等的情况，也就是计算时需要判空

返回一个表示和的链表，直接在较长的链表上操作   创建新链更方便
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val:0}
	n1,n2,carry,current := 0,0,0,head

	for l1!=nil || l2!=nil || carry!=0 {
		if l1==nil {
			n1=0
		}else{
			n1=l1.Val
			l1 = l1.Next
		}
		if l2==nil {
			n2=0
		}else{
			n2=l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val:(n1+n2+carry)%10}
		current = current.Next
		//进位
		carry = (n1+n2+carry)/10	
	}
	return head.Next
}

