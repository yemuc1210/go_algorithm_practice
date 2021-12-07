package NC5
import . "structs"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
  * 根节点到叶子节点得所有路径和  中等
  * @param root TreeNode类 
  * @return int整型
*/
func sumNumbers( root *TreeNode ) int {
    // write code here
	var res int
	dfs(root,0,&res)
	return res
}

func dfs(root *TreeNode,path int,res *int) {
	if root == nil {
		return 
	}
	if root.Left==nil && root.Right==nil{
		path = path*10 +root.Val
		*res += path
		return 
	}
	//访问当前节点
	path = path*10 +root.Val
	//左
	if root.Left != nil {
		dfs(root.Left, path, res)
	}
	if root.Right != nil {
		dfs(root.Right, path, res)
	}
}