//原题 53 最大子序和
/*
给定一个整数数组 nums ，
找到一个具有最大和的连续子数组（子数组最少包含一个元素），
返回其最大和。

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/
package dsday01

//感觉可以用滑动窗口  不过时间复杂度是O(n^2)
/*
动态规划  dp[i]表示[0,i]区间内各个子区间最大的值
状态转移方程：
    dp[i]=nums[i]+dp[i-1]       dp[i-1]>0
	dp[i]=nums[i]    dp[i-1]<=0
*/
func maxSubArray(nums []int)int{
	if len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}

	dp := make([]int,len(nums))
	dp[0]=nums[0]
	res := dp[0]
	for i:=1;i<len(nums);i++ {
		if dp[i-1] > 0{
			dp[i] = nums[i] + dp[i-1]
		}else{
			dp[i] = nums[i]
		}
		if res < dp[i]{
			res = dp[i]
		}
	}
	return res
}