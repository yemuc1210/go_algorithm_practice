package lt279

import "math"

/*
279. 完全平方数
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
*/
/*
1 dfs即可
2 数学
四平方定理：任何一个正整数都可以表示成不超过四个整数的平方和
推论：满足四数平方和定理的数n 必定满足n=4^a*(8b+7)
*/
func numSquares(n int) int {
	for n%4==0 {
		n/=4
	}
	if n%8 == 7 {
		return 4  //推论 满足四数和
	}
	a := 0
	for a*a <= n {
		b := int(math.Sqrt(float64(n-a*a)))
		if a*a + b*b == n {
			if a!=0 && b!= 0{
				return 2
			}
			if a!=0 || b!=0 {
				return 1
			}
			return 0
		}
		a += 1
	}
	return 3
}

/*
dp[i]  最少需要i个数的平方和
这些数必定在区块[1,n**0.5] 可以枚举
	枚举j  然后问题变为求和i-j^2的枚举
	子问题和原问题相似 可以使用动态规划
状态转移方程：
	dp[i]=1 + min{dp[i-j^2]} j∈[1,i**0.5]
	dp[0] 边界
*/
func numSquares1(n int)int{
	dp := make([]int, n+1)
	for i:=1;i<=n;i++{
		minn := math.MaxInt32
		for j:=1;j*j <= i; j++{
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	} 
	return dp[n]
}

func min(a,b int) int{
	if a<b {
		return a
	}
	return b
}