package tencent50
/*
53. 最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），
返回其最大和。
*/
/*
示例 1：
	输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出：6
	解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

算法：滑动窗口
遇到负数，则对当前窗口值是不利的，所以   左右边界如何滑动
	curMax

动态规划，转移方程
	dp[i]: nums[0:i]中最大子序和
	dp[i] = nums[i] + dp[i-1]  dp[i-1]>0
	dppi = nums[i]  dp[i-1] <0
dp[i-1]>0  则以后的元素就可能有贡献；一旦是负的，后面的元素就浪费了
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	dp := make([]int,len(nums))
	dp[0] = nums[0]
	curMax := dp[0]
	for i:=1;i<len(nums);i++{
		if dp[i-1]>0{
			dp[i] = nums[i] + dp[i-1]
		}else{
			dp[i] = nums[i]
		}
		if curMax < dp[i] {
			curMax = dp[i]
		}
	}
	return curMax
}