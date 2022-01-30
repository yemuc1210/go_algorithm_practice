package NC50
import . "structs"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
  * 链表中节点每k个翻转一次  lt25
  * @param head ListNode类 
  * @param k int整型 
  * @return ListNode类
*/
func reverseKGroup( head *ListNode ,  k int ) *ListNode {
    // write code here
  node := head
  for i := 0; i < k; i++ {
    if node == nil {
      return head
    }
    node = node.Next
  }
  newHead := reverse(head, node)
  head.Next = reverseKGroup(node, k)
  return newHead
}
  
func reverse(first *ListNode, last *ListNode) *ListNode {
  prev := last
  for first != last {
    tmp := first.Next
    first.Next = prev
    prev = first
    first = tmp
  }
  return prev
}
  