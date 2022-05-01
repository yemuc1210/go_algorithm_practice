package lt322

import "fmt"

/*
算法计划   dp
322. 零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。



f(i):组成金额i的最小数量
在计算f(i)之前，已经有f(0)~f(i-1)的结果，可以得到转移方程
f(i)=minf(i-cj) + 1     cj：第j枚金币面值    j:0....n-1


*/
func coinChange(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	// max := amount + 1
	dp := make([]int, amount+1)
	//初始化 
	for i := range dp {
		dp[i] = amount+1
	}
	dp[0] = 0

	for i:=1; i<=amount; i++ {
		for j:=0; j< len(coins); j++ {
			if coins[j] <= i { ///不然i-cj没意义
				dp[i] = min(dp[i],dp[i-coins[j]] + 1)
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}