package offerS29
type Node struct {
	     Val int
	     Next *Node
	 }
/*剑指 Offer II 029. 排序的循环链表
给定循环升序列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的。

给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。

如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。

如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。

*/
func insert(aNode *Node, x int) *Node {
	nNode := &Node{Val: x}
	if aNode == nil {
		nNode.Next = nNode
		return nNode
	}
	if aNode == aNode.Next {
		aNode.Next = nNode
		nNode.Next = aNode
		return aNode
	}
	cur := aNode.Next
	insertIdx := aNode
	maxIdx := aNode
	for cur != aNode {
		if cur.Val >= maxIdx.Val {
			maxIdx = cur
		}
		if cur.Val <= x && (insertIdx.Val > x || cur.Val >= insertIdx.Val) {
			insertIdx = cur
		}
		cur = cur.Next
	}
	if maxIdx.Val <= x {
		nNode.Next = maxIdx.Next
		maxIdx.Next = nNode
	} else if insertIdx.Val > x {
		nNode.Next = maxIdx.Next
		maxIdx.Next = nNode
	} else if insertIdx.Val <= x {
		nNode.Next = insertIdx.Next
		insertIdx.Next = nNode
	}
	return aNode
}

