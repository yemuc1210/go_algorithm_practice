package main

import (
	"fmt"
	"strconv"
	"strings"
)
//offerS26
//对于一个链表 L: L0→L1→…→Ln-1→Ln,
// 将其翻转成 L0→Ln→L1→Ln-1→L2→Ln-2→…
// 输入是一串数字，请将其转换成单链表格式之后，再进行操作
type ListNode struct{
	Val int 
	Next *ListNode
}

func main() {
	var line string
	fmt.Scanln(&line)
	a := strings.Split(line,",")
	// fmt.Printf("%#v\n",line)
	// fmt.Println(a,len(a))
	// nums := []int{}
	head := &ListNode{Val: -999999}
	p := head
	for i:=0;i<len(a);i++{
		num,err := strconv.Atoi(a[i])
		if err!=nil {
			panic(err)
		}
		// nums = append(nums, num)
		// fmt.Println(num)
		p.Next = &ListNode{Val: num}
		p = p.Next
	}
	// print(head)
	head = head.Next
	// fmt.Println()
	//中间分 后半部分翻转  然后合并
	//寻找中间节点
	reorderList(head)
	print(head)
}
func print(head *ListNode) {
	for p:=head;p!=nil;p=p.Next {
		if p.Next==nil {
			fmt.Print(p.Val)
		}else{
			fmt.Printf("%v,",p.Val)
		}
	}
}
func reorderList(head *ListNode)  {
	//双指针、首 中间节点
	if head == nil || head.Next == nil {
		return 
	}

	// 寻找中间结点
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	// 反转链表后半部分  1->2->3->4->5->6 to 1->2->3->6->5->4
	preMiddle := p1
	preCurrent := p1.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

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