package NC126

/**
 * 最少货币数  dp  lt322
 * @param arr int整型一维数组 the array
 * @param aim int整型 the target
 * @return int整型
*/
/*
题目描述：
给定数组arr，arr中所有的值都为正整数且不重复。
每个值代表一种面值的货币，每种面值的货币可以使用任意张，再给定一个aim，代表要找的钱数，
求组成aim的最少货币数。
如果无解，请返回-1.

dp[i]  组成金额i的最小硬币数
dp[i] = min{dp[i-cj]}+1  cj是货币面值
*/
func minMoney( arr []int ,  aim int ) int {
	if aim<1 && len(arr)<1 {
		return -1
	}
	dp:= make([]int, aim+1)   //dp[i]  组成金额i的最小硬币数
	//初始化
	for i := range dp{
		dp[i] = aim + 1 //先初始化一个最大值（上界）
	}
	dp[0] = 0
	for i:=1;i<=aim;i++{
		//得到min dp[i-cj]
		for j:=0;j<len(arr);j++{
			if arr[j] <= i {  //货币面值比需要计算的金额小才有意义  不然i-cj<0
				dp[i] = min(dp[i],  dp[i-arr[j]] + 1)
			}
		}
	}
	if dp[aim] > aim {
		return -1
	}
	return dp[aim]
}
func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}