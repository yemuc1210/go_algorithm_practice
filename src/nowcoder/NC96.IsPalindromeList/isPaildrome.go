package NC96
import . "structs"
/**
 * 判断链表是否是回文结构
 * @param head ListNode类 the head
 * @return bool布尔型
*/
func isPail( head *ListNode ) bool {
    // 链表逆置即可 
	if head == nil || head.Next == nil{
		return true
	}
	//得到正序的结果
	arr1 := []int{}
	for p:=head; p!=nil; p = p.Next {
		arr1 = append(arr1, p.Val)
	}
	//逆置链表 逆置使用头插法
	newHead := &ListNode{}
	//可以确保链表至少拥有一个元素
	for p:=head; p!=nil;  {
		//头插到newHead.next
		tmpNode := p.Next
		p.Next = newHead.Next
		newHead.Next = p
		p = tmpNode
	}
	//得到逆置的结果
	arr2 := []int{}
	for p:=newHead.Next; p!=nil;p=p.Next{
		arr2 = append(arr2, p.Val)
	}
	//比较
	if len(arr1) != len(arr2) {
		return false
	}
	for i:=0;i<len(arr1);i++{
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}