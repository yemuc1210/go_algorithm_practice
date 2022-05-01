package lt543
import . "go_practice/structs"
/*
543. 二叉树的直径
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
*/
func diameterOfBinaryTree(root *TreeNode) int {
	res := 1
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int{
		if node==nil{
			return 0
		}
		L := dfs(node.Left)
		R := dfs(node.Right)
		res = max(res, L+R+1)

		return max(L,R) + 1
	}
	dfs(root)
	return res-1
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}