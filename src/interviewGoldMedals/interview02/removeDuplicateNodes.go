package interview02

/*
面试题 02.01. 移除重复节点
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点

进阶：
如果不得使用临时缓冲区，该怎么解决？
*/


type ListNode struct{
	Val int
	Next *ListNode
}
/*
使用哈希表   存储出现的节点
头结点一定不会被删除   枚举待删除结点的前驱节点

*/
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//当前遍历指针pos
	pos := head
	//哈希表
	occurred := map[int]bool { head.Val:true}   //头节点一定不会被移除
	for pos.Next != nil {
		cur := pos.Next
		if !occurred[cur.Val] {
			occurred[cur.Val] = true
			pos = pos.Next
		}else{
			pos.Next = pos.Next.Next  //跳过，也就是移除了  自动回收
		}
	}
	pos.Next = nil
	return head
}