/*
泰波那契序列 Tn 定义如下： 
T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
示例 1：

输入：n = 4
输出：4
解释：
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
*/
package dpDay01
//双100%  0ms   1.9MB
func tribonacci(n int) int {
	if n==0{
		return 0
	}
	if n==2 || n==1{
		return 1
	}
	// //下面对应处理n>=3
	// dp := make([]int,n+1)
	// dp[0],dp[1],dp[2]=0,1,1
	// for i:=3;i<=n;i++{
	// 	dp[i] = dp[i-1]+dp[i-2]+dp[i-3]
	// }
	// return dp[n]
	dp0,dp1,dp2,dp3 := 0,1,1,2
	for i:=3;i<=n;i++{
		dp3 = dp0 + dp1 + dp2
		dp0,dp1,dp2 = dp1,dp2,dp3
	}
	return dp3
}