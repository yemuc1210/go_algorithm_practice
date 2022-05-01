package offer24
import "go_practice/structs"
type ListNode=structs.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//100%   68.58%
func reverseList(head *ListNode) *ListNode {
	//原地逆置即可
	if head==nil || head.Next==nil{
		return head
	}
	//现在起码有两个元素
	var former *structs.ListNode
	cur :=head
	for cur!=nil{
		latter := cur.Next
		cur.Next=former
		former=cur
		cur=latter	
	}
	return former
}