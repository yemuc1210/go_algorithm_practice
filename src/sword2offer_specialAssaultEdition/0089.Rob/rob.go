package offerS89
/*dp  lt198
剑指 Offer II 089. 房屋偷盗
一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
影响小偷偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算 不触动警报装置的情况下 ，
一夜之内能够偷窃到的最高金额。
*/
func rob(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	n := len(nums)
	//dp[i] 表示nums[0...i]的最大值
	dp := make([]int,n+1)
	dp[0],dp[1] = nums[0], max(nums[1],nums[0])

	for i:=2;i<n;i++{
		dp[i] = max(dp[i-1],dp[i-2]+nums[i])
	}
	return dp[n-1]
}
func max(a,b int)int{
	if a>b {
		return a
	}
	return b
}

//存储优化
func rob1(nums []int) int{
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	n := len(nums)
	//dp[i] 表示nums[0...i]的最大值
	dp := make([]int,n+1)
	dp[0],dp[1] = nums[0], max(nums[1],nums[0])
	//只需要两个变量
	curMax,preMax := 0,0
	for i:=0;i<n;i++{
		tmp := curMax
		curMax = max(curMax, nums[i] + preMax)
		preMax = tmp
	}
	return curMax
}