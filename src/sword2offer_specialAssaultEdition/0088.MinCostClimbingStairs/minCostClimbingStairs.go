package offerS88

/*dp  offers88
剑指 Offer II 088. 爬楼梯的最少成本
数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。

每当爬上一个阶梯都要花费对应的体力值，一旦支付了相应的体力值，就可以选择向上爬一个阶梯或者爬两个阶梯。

请找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。


*/
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1{
		return cost[0]
	}
	n := len(cost)
	dp := make([]int,n+1) //dp[i]表示爬到阶梯i的最小代价
	dp[0] = 0
	dp[1] = 0
	//可以选择爬一个或两个   走完最后一个才算
	for i:=2;i<=n;i++{
		dp[i] = min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}
