package lt931

// import "math"

/*
下降路径最小和
给定方形n*n矩阵，找出并返回通过matrix的下降路径的最小和

下降路径：可以从第一行中的任何元素开始，并从每一行中选择一个元素。
	在下一行选择的元素和当前行所选元素最多相隔一列（
	即位于正下方或者沿对角线向左或者向右的第一个元素）。


dp练习
*/
func minFallingPathSum(matrix [][]int) int {
	if len(matrix) < 1{
		return 0
	}
	n := len(matrix)
	dp := make([][]int,n)
	for i := range dp{
		dp[i] = make([]int, n)
	}
	//dp[i][j]表示以（i,j）为终点的路径中的最小值
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if i==0{
				dp[i][j] = matrix[i][j]   //第一行初始化
			}else{
				//其它行
				if j==0{
					//第一列
					dp[i][j] = min(dp[i-1][j]+matrix[i][j], dp[i-1][j+1]+matrix[i][j])
				}else if j==n-1{
					dp[i][j] = min(dp[i-1][j]+matrix[i][j],dp[i-1][j-1]+matrix[i][j])
				}else{
					dp[i][j]=min(min(dp[i-1][j]+matrix[i][j],dp[i-1][j-1]+matrix[i][j]),dp[i-1][j+1]+matrix[i][j])
				}
				// dp[i][j] = max(dp[i-1][j]+matrix[i][j],dp[i-1][j-1])
			}
		}
	}
	min := dp[n-1][0]
	for i:=1;i<n;i++{
		if dp[n-1][i] < min{
			min = dp[n-1][i]
		}
	}
	return min
}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}