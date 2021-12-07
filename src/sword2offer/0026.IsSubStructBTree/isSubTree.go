package offer26
import (

	"structs"
)
type TreeNode=structs.TreeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A==nil && B==nil{
		return true
	}
	if A==nil || B==nil{
		return false
	}
	var res bool
	if A.Val == B.Val{
		res = judge(A,B)
	}

	if !res{
		res = isSubStructure(A.Left,B)
	}
	if !res{
		res = isSubStructure(A.Right,B)
	}
	return res
}
//helper 校验 B 是否与 A 的一个子树拥有相同的结构和节点值
func judge(A,B *TreeNode)bool{
	if B==nil{
		return true
	}
	if A==nil{
		return false
	}
	if A.Val != B.Val{
		return false
	}
	return judge(A.Left,B.Left) && judge(A.Right,B.Right)
}