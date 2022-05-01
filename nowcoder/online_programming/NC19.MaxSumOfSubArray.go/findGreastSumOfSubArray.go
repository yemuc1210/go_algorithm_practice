package NC19

import "math"

/**
 * 连续子数组的最大和  dp即可  lt53
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray( array []int ) int {
    // write code here 正数才会有贡献 负数not ok
	if len(array) == 0 {
		return math.MinInt32
	}
	dp := make([]int, len(array)+1)  //习惯性+1
	/*
	动态规划  dp[i]表示[0,i]区间内各个子区间最大的值
	状态转移方程：
	    dp[i]=nums[i]+dp[i-1]       dp[i-1]>0
		dp[i]=nums[i]    dp[i-1]<=0
	*/
	dp[0] = array[0]
	res := dp[0]
	for i:=1;i<len(array);i++{
		if dp[i-1]>0 {
			dp[i] = dp[i-1] + array[i]
		}else{
			dp[i] = array[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}