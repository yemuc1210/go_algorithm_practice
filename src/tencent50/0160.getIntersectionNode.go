package tencent50
/*  简单
160. 相交链表
*/
/*
常规题
获得长度l1,l2  长的先行abs(l1,l2) 然后两个同步
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA==nil || headB==nil {
		return nil
	}
	len1,len2 := 0,0
	tmp1,tmp2 := headA,headB  //用来遍历两个链表
	for tmp1!=nil || tmp2!=nil {
		if tmp1!=nil{
			tmp1 = tmp1.Next
			len1++
		}
		if tmp2!=nil{
			tmp2 = tmp2.Next
			len2++
		}
	}

	//计算差值
	diff := difference(len1,len2)
	//指针恢复
	tmp1,tmp2 = headA,headB
	//长的先行
	if len1>len2 {
		for diff>0 {
			tmp1 = tmp1.Next
			diff--
		}
	}else{
		for diff>0{
			tmp2 = tmp2.Next
			diff--
		}
	}
	//同步
	for tmp1!=tmp2{
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}
	return tmp1
}

func difference(a,b int) int {
	if a>b{
		return a-b
	}
	return b-a
}
