package main

import (
	// "bufio"
	"fmt"
	// "os"
)

// 字符串s1 s2 允许插入 删除操作
// 令s1==s2的最小操作步数
func main() {
	var s1,s2 string
	// in := bufio.NewScanner(os.Stdin)
	// in.Scan()
	// s1 = in.Text()
	// in.Scan()
	// s2 = in.Text()
	// // 
	s1,s2 = "https://trendmicro.cn","https://trendmicro.com"
	m,n := len(s1),len(s2)
	dp := make([][]int, m+1)
	for i:=range dp{
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i:=0;i<=n;i++{
		// 第一行
		dp[0][i] = i
	}
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			if s1[i-1] == s2[j-1] {
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1])
			}else{
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+2)
			}
		}
	}
	fmt.Println(dp[m][n])
}
func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}