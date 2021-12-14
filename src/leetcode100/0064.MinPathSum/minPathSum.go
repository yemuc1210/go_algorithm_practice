package lt64
/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格 grid ，
请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
*/
/*
dp[i][j]  左上角到【i，j]最小和
*/
func minPathSum(grid [][]int) int {
	if len(grid)==0 {
		return 0
	}
	m,n := len(grid),len(grid[0])
	dp := make([][]int,m)
	for i:=range dp{
		dp[i] = make([]int,n)
		//初始化第一行
		if i==0 {
			dp[i][0] = grid[i][0]
		}else{
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	//初始化第一列
	for j:=1;j<n;j++{ //（0，0）已经初始化
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a,b int) int{
	if a<b {
		return a
	}
	return b
}