package offerS50
import "go_practice/structs"
type TreeNode = structs.TreeNode

/*
剑指 Offer II 050. 向下的路径节点之和
给定一个二叉树的根节点 root ，和一个整数 targetSum ，
求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var rootSum func(*TreeNode, int) int 
	rootSum = func(root *TreeNode, targetSum int) int{
		//以root为根
		if root == nil {
			return 0
		}
		res:=0
		val := root.Val
		if val == targetSum {
			res ++
		}
		res += rootSum(root.Left, targetSum-val)
		res += rootSum(root.Right, targetSum-val)
		return res
	}
	res := rootSum(root,targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}