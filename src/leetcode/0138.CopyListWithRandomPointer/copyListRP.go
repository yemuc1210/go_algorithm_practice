package lt138

/*
复制带随机指针的链表
中等题

*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// Node define
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	tempHead := copyNodeToLinkedList(head)
	return splitLinkedList(tempHead)
}

func splitLinkedList(head *Node) *Node {
	cur := head
	head = head.Next
	for cur != nil && cur.Next != nil {
		cur.Next, cur = cur.Next.Next, cur.Next
	}
	return head
}

func copyNodeToLinkedList(head *Node) *Node {
	cur := head
	for cur != nil {
		node := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next, cur = node, cur.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	return head
}

// public Node copyRandomList(Node head) {
// 	if(head == null){
// 		return null;
// 	}
// 	Node cur = head;
// 	HashMap<Node,Node> map = new HashMap<>();
// 	while(cur!=null){
// 		map.put(cur,new Node(cur.val));
// 		cur = cur.next;
// 	}
// 	cur=head;
// 	while(cur!=null){
// 		map.get(cur).next=map.get(cur.next);
// 		map.get(cur).random=map.get(cur.random);
// 		cur=cur.next;
// 	}
// 	return map.get(head);
// }