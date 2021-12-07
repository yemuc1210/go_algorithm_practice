package NC8
import . "structs"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 二叉树中和为某一值的路径 找出所有路径 PathSum
 * @param root TreeNode类 
 * @param expectNumber int整型 
 * @return int整型二维数组
*/
func FindPath( root *TreeNode ,  expectNumber int ) [][]int {
    // write code here
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	dfs(root, expectNumber, []int{}, &res)
	return res
}
func dfs(root *TreeNode, target int,path []int, res *[][]int) {
	if root == nil {
		return 
	}
	target -= root.Val
	path = append(path, root.Val)
	if target==0 &&root.Left==nil && root.Right==nil{
		//得到一个路径 根到叶子节点
		*res = append(*res, append([]int(nil), path...))
		path = path[:len(path)-1]
	}
	dfs(root.Left,target,path,res)
	dfs(root.Right,target,path,res)
}
