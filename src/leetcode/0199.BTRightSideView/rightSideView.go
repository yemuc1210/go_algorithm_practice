package lt199
import "structs"
type TreeNode = structs.TreeNode
//BFS 记录每层最后一个
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue)>0 {
		size := len(queue)    //每层的宽度
		for i:=0;i<size;i++ {
			node := queue[0]   //出队
			queue = queue[1:]
			if node.Left != nil{
				queue = append(queue, node.Left)
			}
			if node.Right != nil{
				queue = append(queue, node.Right)
			}
			if i== size-1 {
				res = append(res, node.Val)
			}
		}
	}
	return res
}