package lt1480

/*
每日一题：一维数组的动态和   简单

给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。

请返回 nums 的动态和。

*/
func runningSum(nums []int) []int {
	if len(nums) == 0{
		return []int{}
	}
	res := []int{nums[0]}

	curSum := nums[0]
	for i:=1;i<len(nums);i++{
		curSum += nums[i]
		res = append(res, curSum)
	}
	return res
}

func runningSum1(nums []int)[]int{
	for i:=1;i<len(nums);i++{
		nums[i] += nums[i-1]
	}
	return nums
}