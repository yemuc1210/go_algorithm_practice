package NC9
import . "structs"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
  * 二叉树中和为某一值的路径(一) pathSum 给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径。
  * @param root TreeNode类 
  * @param sum int整型 
  * @return bool布尔型
*/
func hasPathSum( root *TreeNode ,  sum int ) bool {
    // write code here  lt112
	if root==nil {
		return false
	}
	if root.Left==nil && root.Right==nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}