package NC13
import . "structs"

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
  * 二叉树的最大深度
  * @param root TreeNode类 
  * @return int整型
*/
func maxDepth( root *TreeNode ) int {
    // write code here
	return helper(root)
}

func helper(p *TreeNode) int{
	if p==nil {
		return 0
	}
	if p!=nil && p.Left==nil && p.Right == nil {
		return 1
	}
	return max( helper(p.Left),helper(p.Right) ) + 1
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}