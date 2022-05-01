package interview02

/*
面试题 02.07. 链表相交

给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

图示两个链表在节点 c1 开始相交：

说明链表最后一段是公共的

题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA==nil || headB==nil{
		return nil
	}
	//首先计算长度
	lengthA,lengthB := 0,0 
	curA,curB := headA,headB
	for curA != nil {
		lengthA ++
		curA = curA.Next
	}
	for curB != nil {
		lengthB ++
		curB = curB.Next
	}

	curA,curB = headA,headB
	//让A是长的
	if lengthB > lengthA {
		lengthA,lengthB = lengthB,lengthA
		curA,curB = curB,curA
	}
	gap := lengthA - lengthB
	for gap > 0 {
		curA = curA.Next
		gap--
	}
	for curA != nil {
		if curA == curB{
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}
