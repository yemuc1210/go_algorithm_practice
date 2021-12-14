package lt152



/*
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组
	（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

动态规划法练习
*/
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0{
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	minimum, maximum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum
		}
		maximum = max(nums[i], maximum*nums[i])
		minimum = min(nums[i], minimum*nums[i])
		res = max(res, maximum)
	}
	return res
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}