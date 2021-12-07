package lt729

type Node struct {
	left, right *Node
	start, end int
}

func (n *Node) insert(node *Node) bool {
   if node.start >= n.end {
	   if n.right == nil {
		   n.right = node
		   return true
	   }
	   return n.right.insert(node)
   }
   if node.end <= n.start {
	   if n.left == nil {
		   n.left = node
		   return true
	   }
	   return n.left.insert(node)
   }
   return false
}

type MyCalendar struct {
   root *Node
}

func Constructor() MyCalendar {
   return MyCalendar{}
}

func (mc *MyCalendar) Book(start int, end int) bool {
   node := &Node{start:start, end:end}
   if mc.root == nil {
	   mc.root = node
	   return true
   }
   return mc.root.insert(node)
}

