package offer06

/**
 * Definition for singly-linked list.*/
import "go_practice/structs"
type ListNode = structs.ListNode
//给定链表头结点，从头到尾反过来每个结点值，数组返回
func reversePrint(head *ListNode) []int {
	//栈
	res := []int{}
	tmp := []int{}
	top := 0  //栈头
	for head != nil {
		tmp = append(tmp, head.Val)
		top+=1
		head = head.Next
	}
	for i:=top-1;i>=0;i--{
		res = append(res, tmp[i])
	}
	return res
}
func reversePrint1(head *ListNode) ([]int) {
	if head == nil {return nil}
	return append(reversePrint(head.Next), head.Val)
}

/*
执行用时：4 ms, 在所有 Go 提交中击败了60.84%的用户
内存消耗：3.5 MB, 在所有 Go 提交中击败了49.97%的用户
*/