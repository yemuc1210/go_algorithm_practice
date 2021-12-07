package tencent50
/*困难
23. 合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/
/*
两两合并
一列一列来，设立k个指针 每次从k个指针中选一个最小的   开销也很大
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	if len(lists) == 1{
		return lists[0]
	}
	k := len(lists) // 
	head := &ListNode{}
	head.Next = lists[0]
	for i:=1;i<k;i++{
		//将head与lists[i]合并
		head.Next = mergeTwo(head.Next,lists[i])
	}

	return head.Next
}

func mergeTwo(l1 *ListNode, l2 *ListNode) *ListNode{
	head := &ListNode{}
	cur := head
	for l1!=nil || l2!=nil {
		if l1!=nil && l2!=nil {
			//较小的插入
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			}else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		}else if l2==nil { //此分支 l1!=nil  l2==nil
			cur.Next =l1
			break  //
		}else{
			cur.Next = l2
			break
		}
	}
	return head.Next
}