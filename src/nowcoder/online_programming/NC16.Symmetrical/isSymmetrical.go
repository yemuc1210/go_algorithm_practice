package NC16
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
 * 对此二叉树  判断是否是自身的镜像
 * @param pRoot TreeNode类 
 * @return bool布尔型
*/
func isSymmetrical( pRoot *TreeNode ) bool {
    // write code here
	return check(pRoot,pRoot)
}

func check(p,q*TreeNode) bool {
	if p==nil && q==nil {
		return true
	}
	if p==nil || q==nil {
		return false  //只有一个为nil
	}
	return p.Val==q.Val && check(p.Left,q.Right) && check(p.Right,q.Left)
}