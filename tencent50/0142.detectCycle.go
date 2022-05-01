package tencent50
/*中
142. 环形链表 II
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。
*/

/*
对链表节点建立  序号-节点  映射表
思路1：根据141的环检测方法 查表，得到节点的序号
	好像不可行，快慢指针相遇的点，不一定是环的初始节点 ??
	答：
分析一下判断环的原理。fast 指针一次都 2 步，slow 指针一次走 1 步。
令链表 head 到环的一个点需要 x1 步，从环的第一个点到相遇点需要 x2 步，
从环中相遇点回到环的第一个点需要 x3 步。那么环的总长度是 x2 + x3 步。

fast 和 slow 会相遇，说明他们走的时间是相同的，可以知道他们走的路程有以下的关系：
```c
fast 的 t = (x1 + x2 + x3 + x2) / 2
slow 的 t = (x1 + x2) / 1

x1 + x2 + x3 + x2 = 2 * (x1 + x2)

所以 x1 = x3
```

所以 2 个指针相遇以后，如果 slow 继续往前走，fast 指针回到起点 head，两者都每次走一步，
那么必定会在环的起点相遇，相遇以后输出这个点即是结果。	

4ms  3.7MB
*/

func detectCycle1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
			return nil
	}
	isCycle, slow := hasCycle142(head)
	if !isCycle {
			return nil
	}
	fast := head
	for fast != slow {
			fast = fast.Next
			slow = slow.Next
	}
	return fast
}
func hasCycle142(head *ListNode) (bool, *ListNode) {
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
					return true, slow
			}
	}
	return false, nil
}
/*
思路2：参照图的方法 计算入度和出度  环的起始节点入度为2（唯一一个）
	入度如何计算
	if p.next!=nil && degree[p.next]==0 {
		degree[p.next] = 1
	}else if !nil && ==1 {
		//这就是入度为2 的
		return p.next
	}
	耗时8ms  内存消耗5.2MB
*/
func detectCycle2(head *ListNode) *ListNode {
	//首先判断有没有环
	if !hasCycle(head) {
		return nil
	}
	newHead := &ListNode{Next:head}
	degree := map[*ListNode]int{}
	p := newHead.Next
	degree[p] = 1
	for p.Next !=nil && degree[p.Next] == 0{
		degree[p.Next] = 1
		p = p.Next
	}
	//跳出循环  degree[p.Next]!=0
	return p.Next
}

// func hasCycle(head *ListNode) bool {

// 	newHead := &ListNode{Next:head}  //忽略head nil判断
// 	fast,slow := newHead, newHead
// 	for fast!=nil && slow!=nil && fast.Next != nil {
// 		fast = fast.Next.Next
// 		slow = slow.Next
// 		if fast == slow {
// 			return true
// 		}
// 	}
// 	return false
// }