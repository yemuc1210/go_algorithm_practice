package NC2
import . "go_practice/structs"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * lt143 链表重排  offer26
 * @param head ListNode类 
 * @return void
*/
func reorderList( head *ListNode )  {
    // write code here
	//找到中间节点，获得后半部分节点 然后反转后半部分节点
    if head==nil || head.Next==nil {
        return 
    }
    //获取中间节点
    p1,p2 := head,head
    for p2.Next!= nil && p2.Next.Next!=nil {
        p1 = p1.Next
        p2 = p2.Next.Next
    }
    //获得后半部分并反转
    preMiddle,preCurrent := p1,p1.Next
    for preCurrent.Next != nil { //头插 反转
        current := preCurrent.Next
        preCurrent.Next = current.Next
        current.Next = preMiddle.Next
        preMiddle.Next = current
    }
    //重新拼接
    // 重新拼接链表  1->2->3->6->5->4 to 1->6->2->5->3->4
    p1 = head
    p2 = preMiddle.Next
    for p1 != preMiddle {
        preMiddle.Next = p2.Next
        p2.Next = p1.Next
        p1.Next = p2
        p1 = p2.Next
        p2 = preMiddle.Next
    }
    return


}