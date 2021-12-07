package lcp34

/*
LCP 34. 二叉树染色
小扣有一个根结点为 root 的二叉树模型，初始所有结点均为白色，
可以用蓝色染料给模型结点染色，模型的每个结点有一个 val 价值。
小扣出于美观考虑，希望最后二叉树上每个蓝色相连部分的结点个数不能超过 k 个，
求所有染成蓝色的结点价值总和最大是多少？
*/
import "structs"
type TreeNode = structs.TreeNode

func maxValue(root *TreeNode, k int) int {
	dp := make(map[*TreeNode][]int)
	return dfs(dp, root, k, k)
}

func dfs(dp map[*TreeNode][]int, root *TreeNode, remain, k int) int {
	if root == nil {
		return 0
	}

	if dp[root] == nil {
		dp[root] = make([]int, k+1)
	}

	if dp[root][remain] > 0 {
		return dp[root][remain]
	}

	out := dfs(dp, root.Left, k, k) + dfs(dp, root.Right, k, k) // root 不染色

	for i := 0; i <= remain-1; i++ {
		l1 := dfs(dp, root.Left, i, k)           // left 染色 i
		r1 := dfs(dp, root.Right, remain-1-i, k) // right 染色 remain-1-i
		out = Max(out, root.Val+l1+r1)           // root 染色
	}
	dp[root][remain] = out
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
