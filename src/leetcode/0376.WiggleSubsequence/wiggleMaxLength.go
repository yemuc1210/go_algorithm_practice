package greedy

/*
376. 摆动序列
如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。
第一个差（如果存在的话）可能是正数或负数。
仅有一个元素或者含两个不等元素的序列也视作摆动序列。
例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，
第二个序列是因为它的最后一个差值为零。
子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列 的长度 。
*/
// 最长子序列 贪心
func wiggleMaxLength(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}
	// 差值 连续递增 或者 连续递减 的 可以删除中间的
	// 有点像买卖股票那题了
	// 只需要返回最长子序列的长度，计算在坡中间（不包括峰/谷值）的节点个数
	var res = 1

	// 统计峰值  原数组中的峰值
	curDiff, preDiff := 0, 0 // 当前差值 和 前一对差值
	for i := 0; i < len(nums)-1; i++ {
		// 当前峰值
		curDiff = nums[i+1] - nums[i]
		// 判断峰值
		if (curDiff > 0 && preDiff <= 0) || (preDiff >= 0 && curDiff < 0) {
			res++
			preDiff = curDiff
		}
	}

	return res
}

func dp(nums []int) int {
	// dp[i][0] 表示考虑前i个数，第i个数作为山峰的摆动子序列的最长长度
	// dp[i][1] 表示考虑前i个数，第i个数作为山谷的摆动子序列的最长长度
	// 推导转移方程
	// dp[i][0] : max(dp[i][0], dp[j][1] + 1)   max(dp[i][0], dp[j][1] + 1)
	// dp[i][1] = max(dp[i][1], dp[j][0] + 1)，表示将nums[i]接到前面某个山峰后面，作为山谷

	// 初始状态
	// dp[0][0] = dp[0][1] = 1
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0],dp[0][1] = 1,1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			}
		}
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
		}
	}
	// 返回
	n := len(nums)
	return max(dp[n-1][0], dp[n-1][1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
