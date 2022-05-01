package offerS8

import "math"

/*
剑指 Offer II 008. 和大于等于 target 的最短子数组
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的
	连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
并返回其长度。如果不存在符合条件的子数组，返回 0 。

递归  dfs?
先dfs找到符合条件的子数组，然后筛选长度最小的
或者设置全局变量存储当前长度最小的值，子数组比该变量小才加入

上述思路：pass  因为要求是连续子数组

思路：滑动窗口

leetcode 209
*/
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		if nums[0] >= target{
			return 1
		}else{
			return 0
		}
	}

	//使用滑动窗口
	left,right := 0,0
	res := math.MaxInt32
	curSum := 0  //当前0

	for left <= right && right < len(nums){
		curSum += nums[right]
		for curSum >= target{  //长度缩减到最小值
			res = min(res, right-left+1) //更新长度
			curSum -= nums[left]
			left ++
		}
		right ++
	}
	if res == math.MaxInt32{
		return 0
	}
	return res
}	

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}