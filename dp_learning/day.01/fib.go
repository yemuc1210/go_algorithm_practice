/*
斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给你 n ，请计算 F(n) 。
*/
package dpDay01
func fib(n int) int {
	if n<0{
		return -1
	}
	if n==0 || n==1 {
		return n
	}
	dp := make([]int,n+1)
	dp[0],dp[1] = 0,1
	for i:=2;i<=n;i++{
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib1(n int) int {
	if n<0{
		return -1
	}
	if n==0 || n==1 {
		return n
	}
	// dp := make([]int,n+1)
	// dp[0],dp[1] = 0,1
	// for i:=2;i<=n;i++{
	// 	dp[i] = dp[i-1] + dp[i-2]
	// }
	// return dp[n]
	//实际上只需要三个量应该就够了
	dp,dp1,dp0 := 1,1,0
	for i:=2;i<=n;i++{
		dp = dp0+dp1
		dp0,dp1 = dp1,dp
	}
	return dp    //空间1.9  75.72%     另外那些神仙用的位运算？
}

// func fib2(n int) int {
// 	f0, f1 := 0, 1

// 	for i := 2; i <= n; i = i + 2 {
// 		f0 = f0 + f1
// 		f1 = f0 + f1
// 	}
// 	if n%2 == 0 {
// 		return f0
// 	}
// 	return f1
// }