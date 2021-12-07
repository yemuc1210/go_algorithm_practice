package offerS54
import "structs"
type TreeNode = structs.TreeNode

/*
给定一个二叉搜索树，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和。
提醒一下，二叉搜索树满足下列约束条件：
	节点的左子树仅包含键 小于 节点键的节点。
	节点的右子树仅包含键 大于 节点键的节点。
	左右子树也必须是二叉搜索树。

要想获得和，得遍历一次
这题怎么也得遍历两次吧
*/
func convertBST(root *TreeNode) *TreeNode {
	//迭代中序遍历
	sum := 0
	stack := []*TreeNode{}
	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			// nodeStack = append(nodeStack, curNode)
			curNode = curNode.Left
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		sum += curNode.Val

		curNode = curNode.Right
	}
	//遍历结束 得到sum  再来一次遍历？
	//栈是空的，不需要重新初始化
	curNode = root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			// nodeStack = append(nodeStack, curNode)
			curNode = curNode.Left
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curVal := curNode.Val
		curNode.Val = sum
		sum -= curVal
		curNode = curNode.Right
	}
	return root
}

//别人的   逆向中序遍历   时间多 空间少   改进下原来的逆中序迭代
func convertBST1(root *TreeNode)*TreeNode{
	sum := 0

	var helper func(*TreeNode) *TreeNode
	helper = func(tn *TreeNode) *TreeNode {
		if tn == nil {
			return tn
		}
		helper(tn.Right)
		tn.Val += sum
		sum = tn.Val
		helper(tn.Left)
		return tn
	}
	helper(root)
	return root
}

func convertBST2(root *TreeNode) *TreeNode {
	//迭代中序遍历
	sum := 0
	stack := []*TreeNode{}
	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			// nodeStack = append(nodeStack, curNode)
			curNode = curNode.Right
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// sum += curNode.Val
		sum += curNode.Val
		curNode.Val = sum

		// curNode = curNode.Right
		curNode = curNode.Left
	}
	
	return root
}