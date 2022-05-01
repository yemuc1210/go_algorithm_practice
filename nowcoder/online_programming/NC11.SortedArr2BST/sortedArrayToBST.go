package NC11
import . "go_practice/structs"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
  * 将升序数组转为平衡二叉搜索树   二分的思路
  * @param num int整型一维数组 
  * @return TreeNode类
*/
func sortedArrayToBST( num []int ) *TreeNode {
    // write code here
	if len(num) == 0 {
		return nil
	}
	return &TreeNode{Val: num[len(num)/2], Left: sortedArrayToBST(num[:len(num)/2]), Right: sortedArrayToBST(num[len(num)/2+1:])}
}