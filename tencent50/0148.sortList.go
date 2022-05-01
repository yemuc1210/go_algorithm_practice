package tencent50
/*
148. 排序链表
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

进阶：
	你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
 
*/

/*
类似归并排序，将链表分为两段进行排序，最后合并  

每两个排序-》4个-》8个。。。

cut(list,n) :分割前n
merge(list1,list2)
*/

func merge(list1,list2 *ListNode) *ListNode {
	newHead := &ListNode{}   //dummy head
	//归并思路 三个数组/链表
	tmp,tmp1,tmp2 := newHead,list1,list2
	for tmp1!=nil && tmp2!=nil {
		if tmp1.Val <= tmp2.Val {
			//插入tmp
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else{
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	//非空判断
	if tmp1 != nil {
		tmp.Next = tmp1
	} else if tmp2 != nil {
		tmp.Next = tmp2
	}
	return newHead.Next

}
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	//对半分  首先需要获得长度
	length := 0
	for node:=head;node!=nil;node=node.Next{
		length++
	}
	newHead := &ListNode{Next: head} // dummyHead
	for subLength:=1; subLength<length; subLength <<= 1 { //1->2->4->8...
		prev,cur := newHead, newHead.Next    
		for cur!=nil {
			head1 := cur  //cut  先切subLength
            for i := 1; i < subLength && cur.Next != nil; i++ {
                cur = cur.Next
            }

            head2 := cur.Next
            cur.Next = nil
            cur = head2
            for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
                cur = cur.Next
            }//cut 结束

            var next *ListNode  //
            if cur != nil {
                next = cur.Next
                cur.Next = nil  //断开待排序 和 未排序部分
            }

            prev.Next = merge(head1, head2)

            for prev.Next != nil {
                prev = prev.Next
            }
			//prev是有序的最后一个
            cur = next
		}

	}
	return newHead.Next
}