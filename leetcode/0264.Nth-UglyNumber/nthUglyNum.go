package lt264
/*
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是只包含质因数 2、3 和/或 5 的正整数。

dp题
dp[i]表示第i个丑数
p2,p3,p5  表示下一个丑数是当前指针指向的丑数乘以对应的质因数，初始为1
当2<=i<=n时  dp[i]=min(dp[p2]*2,dp[p3]*3,dp[p5]*5)
然后比较dp[i]和dp[p2],dp[p3],dp[p5]是否相等，若相等，则对应指针+1
*/


func nthUglyNumber(n int) int {
	dp := make([]int,n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
    for i := 2; i <= n; i++ {
        x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
        dp[i] = min(min(x2, x3), x5)
        if dp[i] == x2 {
            p2++
        }
        if dp[i] == x3 {
            p3++
        }
        if dp[i] == x5 {
            p5++
        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
