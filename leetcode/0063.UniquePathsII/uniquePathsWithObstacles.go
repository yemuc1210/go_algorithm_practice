package lt63
/*中
63. 不同路径 II
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。
*/
/*
依然可以用dp
机器人行走规律没变
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid)==0 || len(obstacleGrid[0])==0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m,n := len(obstacleGrid),len(obstacleGrid[0])  //m行n列
	dp := make([][]int, m)
	for i:=range dp{
		dp[i] = make([]int, n)
		//初始化第一行    dp[i][0] = 1 if dp[i-1][0]==1 && obGrid[i][0]==0
		if i==0 {
			dp[i][0] = 1
		}else{
			if dp[i-1][0]==1 && obstacleGrid[i][0]==0{
				dp[i][0] = 1
			}else{
				dp[i][0] = 0
			}
		}
	}
	//初始化第一行 第一列
	for j:=1;j<n;j++{
		if dp[0][j-1] == 1 && obstacleGrid[0][j]==0 {
			dp[0][j]=1
		}//else dp[0][j]=1
	}
	
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			//dp[i][j] = dp[i-1][j] + dp[i][j-1]
			//判断dp[i][j]需要从哪获得
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}