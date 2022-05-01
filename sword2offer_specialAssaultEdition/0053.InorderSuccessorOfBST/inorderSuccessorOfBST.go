package offerS5

import "go_practice/structs"
type TreeNode = structs.TreeNode

/*
剑指 Offer II 053. 二叉搜索树中的中序后继
给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。
如果节点没有中序后继，请返回 null 。

节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。
*/
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	//迭代中序
	stack := []*TreeNode{}
	found := false
	curNode := root

	for len(stack) > 0 || curNode != nil {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}   //中序遍历  首先把左节点都入栈
		curNode = stack[len(stack)-1]   //然后弹出最后一个左节点
		stack = stack[:len(stack)-1]
		if found {
			break
		}
		if curNode == p {
			found = true
		}
		curNode = curNode.Right
	}
	return curNode
}