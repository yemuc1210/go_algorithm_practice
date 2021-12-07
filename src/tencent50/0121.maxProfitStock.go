package tencent50

/*
121. 买卖股票的最佳时机
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。
设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/
/*
跌买涨卖
dp
栈 与栈顶比较
*/

//单调栈
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	stack, res := []int{prices[0]}, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] { //比栈顶还大 收益+
			stack = append(stack, prices[i])
		} else {
			idx := len(stack) - 1
			for ; idx >= 0; idx-- {
				if stack[idx] < prices[i] {
					break //当前栈顶比price[i]小  说明正收益
				}
			}
			stack = stack[:idx+1]
			stack = append(stack, prices[i])
		}
		res = max(res, stack[len(stack)-1]-stack[0])
	}

	return res
}

// func max(a int, b int) int {
// 	if a > b {
// 			return a
// 	}
// 	return b
// }

// 模拟 DP
func maxProfit1(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxProfit
}
