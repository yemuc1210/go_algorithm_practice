package NC133
import . "go_practice/structs"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 链表奇偶重排 
 * @param head ListNode类 
 * @return ListNode类
*/
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