package lt61
import "structs"
type ListNode = structs.ListNode
/*
61. 旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置
*/
/*
移位的感觉
1.先转换成数组，然后移位，构建新的链表
2.在链表上操作
循环旋转，所以最终状态确定  k-lenght
	每次移动一个链表
	具体操作：得到最后一个节点，添加到head
	问题关键：每次得到最后一个节点，都需要遍历一遍链表
	解决：可以先得到倒数k个节点，然后将这k个节点插入链表头
*/
func rotateRight(head *ListNode, k int) *ListNode {
	//判空
	if head == nil || k==0 || head.Next==nil{
		return head
	}
	newHead := &ListNode{Val: 0,Next: head}
	length := 0
	cur := newHead
	for cur.Next!= nil { //得到长度
		cur = cur.Next
		length ++
	}
	if k%length == 0{  //最终状态是可以确定的
		return head
	}
	cur.Next = head  //将最后一个元素指向head，形成环
	cur = newHead     //然后从head开始，将最后K个断开
	for i:=length - k%length;i>0;i-- {
		cur= cur.Next //取最后k个元素
	}
	res := &ListNode{Val:0,Next: cur.Next} //让后面k个成为链表头
	cur.Next = nil
	return res.Next
}
