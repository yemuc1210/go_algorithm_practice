package main

// import "fmt"

// import . "nc_tools"

  type ListNode struct{
    Val int
    Next *ListNode
  }

type elem struct {
	pre,next *ListNode
}
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 
 * @param a ListNode类一维数组 指向每段碎片的开头
 * @return ListNode类
*/
func solve( a []*ListNode ) *ListNode {
    // write code here
	// 字典序:两个链表,第一个不相同节点的位置小的 则字典序小
	m := make(map[*ListNode]elem)
	for _,a1 := range a {
		cur := a1 
		var pre *ListNode
		for cur!=nil {
			if _,ok := m[cur]; ok {
				continue
			}else{
				// 不存在
				e := elem{pre: pre, next: cur.Next}
				m[cur] = e
			}
			pre = cur
			cur = cur.Next
		}
	}
	// 根据map得到环
	newHead := &ListNode{}
	cnt := 0
	preNode := &ListNode{}
	cur := a[0]
	resNode := &ListNode{}
	for cur != preNode {
		if cnt == 0 {
			if cur.Val == 1 {
				resNode = cur
			}
			e := m[cur]
			newHead.Next = cur
			cur = e.next
			preNode = e.pre
			cnt++
		}else{
			if cur.Val == 1 {
				resNode = cur
			}
			e := m[cur]
			newHead.Next = cur
			cur = e.next
			cnt++
			
		}
	}
	return resNode
}