package offerS91
/*lt256
剑指 Offer II 091. 粉刷房子
假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。

当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的正整数矩阵 costs 来表示的。

例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。

请计算出粉刷完所有房子最少的花费成本。

相邻不同色
*/
func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([]int, 3)
	dp[0] = costs[0][0]
	dp[1] = costs[0][1]
	dp[2] = costs[0][2]
	for i := 1; i < n; i++ {
		tmp := make([]int, 3)
		tmp[0] = min(dp[1], dp[2]) + costs[i][0]
		tmp[1] = min(dp[0], dp[2]) + costs[i][1]
		tmp[2] = min(dp[0], dp[1]) + costs[i][2]
		dp = tmp
	}
	return min(dp[0], min(dp[1], dp[2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
