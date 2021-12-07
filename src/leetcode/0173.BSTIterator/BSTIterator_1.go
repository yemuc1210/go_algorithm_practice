package lt173

/*
先做一次中序遍历，保存到数组中，然后对数组进行遍历
*/
type BSTIterator1 struct{
	arr []int
}

func Constructor1(root *TreeNode)(it BSTIterator1){
	it.inorderTraversal(root)
	return
}
func (it *BSTIterator1)inorderTraversal(node *TreeNode){
	if node == nil{
		return 
	}
	it.inorderTraversal(node.Left)
	it.arr = append(it.arr, node.Val)
	it.inorderTraversal(node.Right)
}
func (it *BSTIterator1)Next()int{
	val := it.arr[0]
	it.arr = it.arr[1:]
	return val
}
func (it *BSTIterator1) HasNext() bool {
    return len(it.arr) > 0
}