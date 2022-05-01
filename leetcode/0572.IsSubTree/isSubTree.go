package lt572
import "go_practice/structs"
type TreeNode = structs.TreeNode

/*

算法20天    dfs/bfs
一个树是另一个树的子树 则

	要么这两个树相等
	要么这个树是左树的子树
	要么这个树hi右树的子树


*/
func isSametree(root *TreeNode, subRoot *TreeNode) bool {
	if root==nil && subRoot==nil{
		return true
	}
	return (root!=nil && subRoot!=nil) && root.Val == subRoot.Val && isSametree(root.Left,subRoot.Left) && isSametree(root.Right,subRoot.Right)
}

func isSubtree(s,t *TreeNode)bool{
	if s==nil && t==nil{
		return true
	}
	if s==nil && t!=nil {
		return false
	}
	return isSametree(s,t) || isSubtree(s.Left,t) || isSubtree(s.Right, t)
}