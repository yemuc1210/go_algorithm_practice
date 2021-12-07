package leetcode


import "structs"

type ListNode = structs.ListNode


func deleteDuplicates(head *ListNode) *ListNode {
	//要删除元素，需要一个辅助
	// p:= head
	// q := p.Next
	// for q != nil{
	// 	for q !=nil && q.Val == p.Val{
	// 		q = q.Next  //如果相等，就一直向后遍历
	// 	}
	// 	//此时q.val != p.val
	// 	p.Next = q    //p指向q
	// 	q = q.Next  
	// }
	// return head

	cur := head
	for cur != nil && cur.Next != nil{
		if cur.Val == cur.Next.Val{
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return head
}