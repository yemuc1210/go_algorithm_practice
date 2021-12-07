package offer60

/*
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。
	输入n，打印出s的所有可能的值出现的概率。
你需要用一个浮点数数组返回答案，
其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

使用动态规划
dp[n][s]:n和骰子，和为s的概率
多一个骰子，其概率和上一行概率有关
*/

func dicesProbability(n int) []float64 {
	dp := make([][]float64, n)
	for i := range dp{
		dp[i] = make([]float64, (i + 1) * 6 - i)
	}
	for i := range dp[0]{
		dp[0][i] = float64(1) / float64(6)
	}
	for i := 1; i < len(dp); i ++{
		for j := range dp[i - 1]{
			for k := range dp[0]{
				dp[i][j + k] += float64(dp[i - 1][j]) * float64(dp[0][k])
			}
		}
	}
	return dp[n - 1]
}
	
