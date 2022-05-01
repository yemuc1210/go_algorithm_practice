package NC98

import . "go_practice/structs"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 判断t1树中是否有与t2树完全相同的子树
 * 进阶：空间复杂度: O(1)，时间复杂度 O(n)
 * offer26 isSubStructure lt572 isSubTree
 * @param root1 TreeNode类 
 * @param root2 TreeNode类 
 * @return bool布尔型
*/
func isContains( root1 *TreeNode ,  root2 *TreeNode ) bool {
    // write code here
	if root1==nil && root2==nil {
		return true
	}
	if root1==nil && root2!=nil {
		return false
	}
	return isSame(root1,root2) || isContains(root1.Left, root2) || isContains(root1.Right,root2)
}

func isSame(src,target *TreeNode) bool{
	if src==nil && target==nil {
		return true
	}
	return (src!=nil&&target!=nil) && (src.Val==target.Val) && isSame(src.Left, target.Left) && isSame(src.Right, target.Right)
}