package lt1567

/*
给你一个整数数组 nums ，请你求出乘积为正数的最长子数组的长度。
一个数组的子数组是由原数组中零个或者更多个连续数字组成的数组。
请你返回乘积为正数的最长子数组长度。
*/
func getMaxLen(nums []int)int{

	//包含nums[i]的最大值
	maxdp:=make([]int,len(nums)+1)
	//包含nums[i]的最小值
	mindp:=make([]int,len(nums)+1)
	max:=0
	// 遍历
	for i:=1;i<=len(nums);i++{
		//如果当前值是0
		//那么包含当前值的最大最小都是0
		if nums[i-1]==0{
			maxdp[i],mindp[i]=0,0
		// 如果当前值大于0
		}else if nums[i-1]>0{
			//最大值等于前一个最大值+1
			maxdp[i]=maxdp[i-1] + 1
			// 如果前一个最小值不等于0 最小值等于前一个加一
			if mindp[i-1]!=0{
				mindp[i]=mindp[i-1]+1
			}
		}else{
			//如果当前值小于0
			// 如果前一个最小值大于0，那么最大值就是前一个最小值+1
			if mindp[i-1]>0{
				maxdp[i]=mindp[i-1]+1
			}
			// 最小值等于 最大值+1
			mindp[i]=maxdp[i-1]+1
		}
		// 记录最大的
		if maxdp[i]>max{
			max=maxdp[i]
		}
	}
	return max
	
}