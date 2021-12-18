package NC34

/**
  * 机器人到达终点的路径数
  * @param m int整型 
  * @param n int整型 
  * @return int整型
*/
func uniquePaths( m int ,  n int ) int {
    dp := make([][]int,m)
	for i:=range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j:=1;j<n;j++{
		//初始化第一列
		dp[0][j]= 1
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}