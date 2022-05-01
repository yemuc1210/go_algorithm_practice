package lt99

import (
	"go_practice/structs"
)
type TreeNode = structs.TreeNode
/*
99. 恢复二叉搜索树
给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。
请在不改变其结构的情况下，恢复这棵树。

进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗？
*/

/*
二叉搜索树的性质：中序遍历是递增有序的
思路1：O(n)的解决方法：先遍历，存储遍历结果，得到逆序对，交换之
思路2：遍历过程中在左子树中如果出现了前一次遍历的结点的值大于此次根节点的值，这就出现了出错结点了，记录下来。
	继续遍历直到找到第二个这样的结点。
	最后交换这两个结点的时候，只是交换他们的值就可以了，而不是交换这两个结点相应的指针指向。
*/
func recoverTree(root *TreeNode)  {
	if root == nil {
		return 
	}
	var prev, target1,target2 *TreeNode
	_,target1,target2 = inorder(root,prev,target1,target2)
	if target1!=nil && target2 !=nil {
		target1.Val,target2.Val = target2.Val,target1.Val
	}
}

func inorder(root,prev,target1,target2 *TreeNode)(*TreeNode,*TreeNode,*TreeNode){
	if root == nil {
		return prev,target1,target2
	}

	prev,target1,target2 = inorder(root.Left, prev,target1,target2)
	if prev!=nil && prev.Val > root.Val {
		if target1 == nil {
			target1 = prev
		}
		target2 = root
	}
	prev = root //prev 等同于当前遍历节点
	prev,target1,target2 = inorder(root.Right, prev,target1,target2)
	return prev,target1,target2
}
