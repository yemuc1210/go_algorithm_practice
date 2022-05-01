package tencent50
/*
62. 不同路径
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
*/

/*
dp
dp[i][j]:从左上角到（i，j）的路径数
状态转移方程：
	dp[i][j] = dp[i-1][j]+dp[i][j-1]
*/
func uniquePaths(m int, n int) int {
	dp := make([][]int,m)
	for i:= range dp{
		dp[i] = make([]int, n)
		//初始化
		dp[i][0] = 1  //第一列
	}
	//第一行初始化
	for j:=0;j<n;j++{
		dp[0][j] = 1  //第一行
	}

	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}