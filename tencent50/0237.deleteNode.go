package tencent50
/*简单
237. 删除链表中的节点
请编写一个函数，用于 删除单链表中某个特定节点 。
在设计函数时需要注意，你无法访问链表的头节点 head ，只能直接访问 要被删除的节点 。

题目数据保证需要删除的节点 不是末尾节点 。
*/
/*
无法访问head 只能访问当前待删除节点
所以复制后一个，然后删除后一个
*/
func deleteNode(node *ListNode) {
    node.Val = node.Next.Val;
    // node.next = node.next.next;
    // *node = *(node.Next)
    node.Next = node.Next.Next
}