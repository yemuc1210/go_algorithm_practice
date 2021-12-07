package offerS90
/*
剑指 Offer II 090. 环形房屋偷盗
一个专业的小偷，计划偷窃一个环形街道上沿街的房屋，
每间房内都藏有一定的现金。
这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组 nums ，
请计算 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

89的加强版，首尾相连
增加额外的转换 取0就不能取n-1  
在【0，n-1]中找最大值  【1，n]中找最大值
*/
func rob(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	if len(nums) == 2{ //首尾相连
		return max(nums[0],nums[1])
	}
	n := len(nums)
	return max(helper(nums,0,n-2), helper(nums,1,n-1))
}
func helper(nums []int, start,end int) int{
	//同89
	preMax := nums[start]
	curMax := max(preMax, nums[start+1])
	for i:=start+2;i<=end;i++{
		tmp:=curMax
		curMax = max(curMax, preMax+nums[i])
		preMax = tmp
	}
	return curMax
}
func max(a,b int)int{
	if a>b {
		return a
	}
	return b
}