package lt328
import . "go_practice/structs"
//奇偶链表重排 中等
func oddEvenList( head *ListNode ) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next==nil{
		return head
	}
    //奇数位节点放一起  偶数位节点放一起  奇前
	oddList,evenList := &ListNode{Next: head},&ListNode{Next: head.Next}
	oddLast,evenLast := oddList.Next, evenList.Next
	
	idx := 3
	cur := evenLast.Next  //第三个  也是主链表
	oddLast.Next = nil
    evenLast.Next = nil  //even断开
	for cur != nil {
        tmp := cur.Next
		if idx%2==0 {
			//插入偶数
// 			tmp := cur.Next
			evenLast.Next = cur
			evenLast = cur
            evenLast.Next = nil
// 			cur = tmp
		}else{
			//奇数
// 			tmp := cur.Next
			oddLast.Next = cur
			oddLast = cur
            oddLast.Next = nil
// 			cur = tmp
		}
        cur = tmp 
        idx ++
	}
	//合并
	oddLast.Next = evenList.Next
	return oddList.Next
}