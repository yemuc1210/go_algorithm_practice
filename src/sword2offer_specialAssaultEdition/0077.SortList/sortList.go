package offerS77

import (
	"sort"
	"structs"
)
type ListNode = structs.ListNode
/*  lt148
剑指 Offer II 077. 链表排序
给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Next: head}
	//得到数组
	arr := []int{}
	for newHead.Next != nil {
		newHead = newHead.Next
		arr = append(arr, newHead.Val)
	}
	arr = append(arr, newHead.Next.Val)
	
	//数组排序
	sort.Ints(arr)

	//重新构建链表
	newHead.Next = head
	idx:=0
	for newHead.Next != nil{
		newHead=newHead.Next
		newHead.Val = arr[idx]
		idx ++
	}
	newHead.Next.Val = arr[idx]
	return head.Next
}

//上面的不算，只是数字排序
func merge(head1, head2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    temp, temp1, temp2 := dummyHead, head1, head2
    for temp1 != nil && temp2 != nil {
        if temp1.Val <= temp2.Val {
            temp.Next = temp1
            temp1 = temp1.Next
        } else {
            temp.Next = temp2
            temp2 = temp2.Next
        }
        temp = temp.Next
    }
    if temp1 != nil {
        temp.Next = temp1
    } else if temp2 != nil {
        temp.Next = temp2
    }
    return dummyHead.Next
}

func sortList1(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    length := 0
    for node := head; node != nil; node = node.Next {
        length++
    }

    dummyHead := &ListNode{Next: head}
    for subLength := 1; subLength < length; subLength <<= 1 {
        prev, cur := dummyHead, dummyHead.Next
        for cur != nil {
            head1 := cur
            for i := 1; i < subLength && cur.Next != nil; i++ {
                cur = cur.Next
            }

            head2 := cur.Next
            cur.Next = nil
            cur = head2
            for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
                cur = cur.Next
            }

            var next *ListNode
            if cur != nil {
                next = cur.Next
                cur.Next = nil
            }

            prev.Next = merge(head1, head2)

            for prev.Next != nil {
                prev = prev.Next
            }
            cur = next
        }
    }
    return dummyHead.Next
}