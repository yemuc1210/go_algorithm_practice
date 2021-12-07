package interview04

/*   leetcode 285
面试题 04.06. 后继者
设计一个算法，找出   二叉搜索树  中指定节点的“下一个”节点（也即中序后继）。

如果指定节点没有对应的“下一个”节点，则返回null。

写递归版本的中序遍历
*/
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
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

//根据二叉搜索树的性质找
func inorderSuccessor1(root,p *TreeNode) *TreeNode{
	curNode, ans := root, (*TreeNode)(nil)
	for curNode != nil {
		//若后继
		if curNode.Val > p.Val {
			if ans == nil || ans.Val > p.Val {
				ans = curNode
			}
		}
		//p
		if curNode.Val == p.Val {
			//如果有右子树，则在右子树
			if curNode.Right != nil {
				//右子树最左
				curNode = curNode.Right
				for curNode.Left != nil {
					curNode = curNode.Left
				}
				//return 
				return curNode
			}
			break
		}
		if curNode.Val > p.Val {
			curNode = curNode.Left  //左子树
		}else{
			curNode = curNode.Right
		}
	}
	return ans
}