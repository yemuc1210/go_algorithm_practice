package NC14

import "go_practice/structs"
type TreeNode = structs.TreeNode
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 锯齿形层序遍历  之字形打印二叉树  lt103
 * @param pRoot TreeNode类 
 * @return int整型二维数组
*/
func Print( pRoot *TreeNode ) [][]int {
    // 记录每层的遍历结果 按顺序选择输出顺序即可
	if pRoot== nil {
		return [][]int{}
	}
	queue := []*TreeNode{pRoot}
	res := [][]int{}
	isOdd := false
	for len(queue)>0 {
		curLevelWidth := len(queue)  //当前层的宽度
		curLevelNodes := []*TreeNode{}
		for i:=0;i<curLevelWidth;i++{
			curNode := queue[0]
			queue = queue[1:] //队头出队
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
			curLevelNodes = append(curLevelNodes, curNode)
		}
		//按照层次序号 选择输出顺序
		if isOdd {  //奇数层 逆序
			//当成栈弹出
			tmp := []int{}
			for i:=len(curLevelNodes)-1;i>=0;i--{
				tmp = append(tmp, curLevelNodes[i].Val)
			}
			//保存到结果
			res = append(res, append([]int(nil), tmp...))
			isOdd = false
		}else {//偶数层 正序 从左-右
			tmp := []int{}
			for i:=0;i<len(curLevelNodes);i++{
				tmp = append(tmp, curLevelNodes[i].Val)
			}
			res = append(res, append([]int(nil), tmp...))
			isOdd = true
		}
	}
	return res
}