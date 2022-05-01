package leetcode
/*

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
*/
/*动态规划
dp[i]表示爬到i级的方案数，那么考虑最后一步，可能有两种情况，最后一步是两级别
	最后是一级
得到状态转移方程  dp[i]=dp[i-1]+dp[i-2]

*/
func climbStairs(n int) int {
	//最简单的几个情况
	if n==0{
		return 0
	}
	
	x2,x1,x := 0,0,1
	for i:=1 ; i<=n; i++ {
		x2 = x1
		x1 = x
		x = x1+x2
	}
	return x 
}