/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package offer22
import "structs"
type ListNode=structs.ListNode
func getKthFromEnd(head *ListNode, k int) *ListNode {
	//双指针
	former,latter:=head,head
	for i:=0;i<k;i++{
		if former==nil{
			return nil
		}
		former=former.Next
	}
	for former!=nil{
		former=former.Next
		latter=latter.Next
	}
	return latter
}