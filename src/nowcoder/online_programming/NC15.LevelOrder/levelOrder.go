package NC15
import . "structs"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
  * 层序遍历
  * @param root TreeNode类 
  * @return int整型二维数组
*/
func levelOrder( root *TreeNode ) [][]int {
    // write code here
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := [][]int{}  //每一层作为一个数组插入
	for len(queue) > 0 {
		levelSize := len(queue)
		curLevelNode := []int{}
		for i:=0;i<levelSize;i++{
			curNode := queue[0]
			queue = queue[1:]
			curLevelNode = append(curLevelNode, curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		//当前层插入
		res = append(res, append([]int(nil),curLevelNode...))
	}
	return res
}