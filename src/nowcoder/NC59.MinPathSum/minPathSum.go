package NC59

/**
 * 矩阵的最小路径和 lt64
 * @param matrix int整型二维数组 the matrix
 * @return int整型
*/
func minPathSum( matrix [][]int ) int {
    // 左上角每次只能向右或向下走，到达右下角最小值  dp
	if len(matrix) == 0 || len(matrix[0])==0 {
		return 0
	}
	m,n := len(matrix),len(matrix[0])  //m*n
	dp := make([][]int,m)
	for i:=range dp{
		dp[i] = make([]int, n)
		//第一列
		if i==0{
			dp[i][0] = matrix[i][0]
		}else{
			dp[i][0] = dp[i-1][0] + matrix[i][0]
		} //第一列只能从上面一个到达
	}
	//第一行
	for j:=1;j<n;j++{
		dp[0][j] = dp[0][j-1] + matrix[0][j]
	}
	//dp
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a,b int) int {
	if a<b{
		return a
	}
	return b
}