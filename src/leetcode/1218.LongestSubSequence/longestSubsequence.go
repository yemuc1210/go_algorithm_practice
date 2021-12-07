package lt1218

/*
1218. 最长定差子序列
给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，
该子序列中相邻元素之间的差等于 difference 。

子序列 是指在不改变其余元素顺序的情况下，
通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。
*/
func longestSubsequence(arr []int, difference int) int {
	res := 0
	if arr == nil{
		return 0
	}
	if len(arr) == 1{
		return 1
	}

	//dp    dp[i] = dp[i-1]   dp = dp[i-1]+1
	/*
	dp[i]以arr[i]结尾的最长等差序列
	左侧找到arr[j] = arr[i]-d，将arr[i]加至arr[j]结尾的最长等差子序列
	递归地从dp[j]计算dp[i]
	相同元素：下标较大的dp>=下标较小的         dp[i] = dp[j]+1
				dp[v] = dp[v-d]+1
	*/
	dp := make(map[int]int)
	for _,v := range arr{
		dp[v] = dp[v-difference] + 1
		if dp[v] > res {
			res = dp[v]
		}
	}


	return res
}