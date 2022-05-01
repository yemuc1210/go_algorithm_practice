package tencent50
/* 简单
26. 删除有序数组中的重复项
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，
使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 
并在使用 O(1) 额外空间的条件下完成。
*/
/*
返回 数组长度。。。。 那没必要修改了吧
题目判定时，会根据长度输出数组，所以数组还是需要修改的
*/
func removeDuplicates(nums []int) int {	
	if len(nums) == 0 {
		return 0
	}
	//数组有序，则不必额外排序
	idx := 0
	last := 0  //结果数组的最后一位
	for last < len(nums)-1 {
		//找到下一个
		for nums[idx] == nums[last] {
			idx ++ 
			if idx == len(nums) {
				//越界  本次循环是最后一个
				return last + 1
			}
		}
		nums[last+1] = nums[idx]
		last ++ 
	}
	return last+1
}