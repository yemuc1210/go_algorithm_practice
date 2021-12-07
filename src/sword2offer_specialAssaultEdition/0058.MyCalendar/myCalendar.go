package offerS58

/*
在插入新日程的时候，如果集合用切片，则有一一向后移动元素的复杂度在，
有没有一个数据结构能在常数时间插入元素，还能保证元素有序，
且能在常数时间查询任意索引元素从而能用二分法做查询呢？
堆、list和切片都不行

二叉搜索树BST可堪此任，不过如果插入的顺序比较特别，bst会退化成一个链表
实际需要一个能维持平衡的搜索树，比如红黑树，像Java有TreeMap可用，
Go标准库并没有这样的数据结构，手动实现起来有些复杂，暂不尝试
仅朴素BST来一把
时间复杂度最坏O(n), 平均情况O(lgn)；空间复杂度O(n)
*/
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



/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */