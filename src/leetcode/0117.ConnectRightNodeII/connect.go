package lt117

/*

算法20天    dfs/bfs
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

 

进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。


 */
 type Node struct {
	Val int
	Left *Node
	Right *Node
		Next *Node
}
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil && root.Right!=nil{
		root.Left.Next=root.Right
	}
	if root.Left!=nil && root.Right==nil{
		root.Left.Next=getNext(root.Next)
	}
	if root.Right!=nil {
		root.Right.Next = getNext(root.Next)
	}

	connect(root.Right)
	connect(root.Left)
	return root
}

func getNext(root *Node) *Node{
	if root ==nil {
		return nil
	}
	if root.Left!=nil{
		return root.Left
	}
	if root.Right!=nil{
		return root.Right
	}
	if root.Next!=nil {
		return getNext(root.Next)
	}
	return nil
}