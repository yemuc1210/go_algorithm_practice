package offerS103
/*lt322  dp
剑指 Offer II 103. 最少的硬币数目
给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为每种硬币的数量是无限的。
*/

/*
dp[i] : 金额i需要的最少硬币个数
dp[0] = 0
dp[1] 

dp[i] = min{dp[i-c_j]+1}
*/
func coinChange(coins []int, amount int) int {
	//边界处理
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	dp := make([]int, amount+1)
	//初始化  因为需要得到最小值 所以先设置一个最大值
	for i:=range dp{
		dp[i] = amount+1
	}
	dp[0] = 0

	for i:=1;i<=amount;i++{
		for j:=0;j<len(coins);j++{
			if coins[j] <= i {  //不然i-cj为负
				dp[i] = min(dp[i],  dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount{
		return -1
	}
	return dp[amount]
}

func min(a,b int)int{
	if a<b {
		return a
	}
	return b
}