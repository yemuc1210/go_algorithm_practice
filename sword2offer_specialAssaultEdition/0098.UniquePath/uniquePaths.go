package offerS98
/* lt62
剑指 Offer II 098. 路径的数目
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

*/
/*
dp[i][j] 表示左上角走到(i,j)的数量
dp[i][j] = dp[i-1][j] + dp[i][j-1]
*/
func uniquePaths(m int, n int) int {
	dp:=make([][]int,m)
	for i:=range dp{
		dp[i] = make([]int, n)
		dp[i][0] = 1  //  左上角来只有一条路 同一列
	}
	for j:=0;j<n;j++{
		dp[0][j] = 1  //第一行
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j] = dp[i-1][j]+dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}