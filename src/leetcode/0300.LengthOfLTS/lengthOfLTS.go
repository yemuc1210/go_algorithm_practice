package lt300
/*
算法计划  dp    673
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


dp[i]  考虑前i个  以i为结尾的最长上升子序列  nums[i]必选

从前往后
dp[i] = max(dp[j]) + 1     nums[j]<nums[i]
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := []int{}
	maxDP := 0
	for i:=0;i<len(nums);i++{
		dp = append(dp, 1)
		for j:=0;j<i;j++{
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxDP{
			maxDP = dp[i]
		}
	}
	return maxDP
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}