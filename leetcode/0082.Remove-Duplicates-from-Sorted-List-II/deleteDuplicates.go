package lt82

// 删除有序链表中的重复元素
func deleteDuplicates22(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    dummy := &ListNode{0, head}

    cur := dummy
    for cur.Next != nil && cur.Next.Next != nil {
        if cur.Next.Val == cur.Next.Next.Val {
            x := cur.Next.Val  // 待删除的元素值
            for cur.Next != nil && cur.Next.Val == x {
                cur.Next = cur.Next.Next  // 删除cur.Next元素
            }
        } else {
            cur = cur.Next
        }
    }

    return dummy.Next
}
 