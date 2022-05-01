package lt148

import "go_practice/structs"

type ListNode = structs.ListNode
/*
148. 排序链表    147
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

进阶：

你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？


两两合并
	merge(list1,list2)
	cut(list,n)   分割前n，返回后半部分链表头
*/
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

func sortList(head *ListNode) *ListNode {
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

