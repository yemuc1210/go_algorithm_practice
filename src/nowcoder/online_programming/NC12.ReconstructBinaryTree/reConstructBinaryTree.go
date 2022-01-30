package NC12
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
 * 重建二叉树 给定前序和中序遍历结果，构建二叉树
 * @param pre int整型一维数组 
 * @param vin int整型一维数组 
 * @return TreeNode类
*/
func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
    /* lt105
	前序规律：第一个为根
	中序：根在中间  分割左右子树
	*/
	//查找得到根在中序数组中的位置  使用map
	inPos := make(map[int]int)
	for i:=0;i<len(vin);i++{
		inPos[vin[i]] = i
	}
	return build(pre,0,len(pre)-1,0,inPos)
}

func build(pre []int, preStart,preEnd int, inStart int, inPos map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: pre[preStart]} //构建根节点
	rootIdx := inPos[root.Val]  //得到中序数组的根节点下标
	leftLen := rootIdx - inStart   //左子树长度
	root.Left = build(pre,preStart+1,preStart+leftLen, inStart,inPos)
	root.Right = build(pre, preStart+leftLen+1, preEnd,rootIdx+1, inPos)
	return root
}
