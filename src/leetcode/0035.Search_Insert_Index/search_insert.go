/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
*/

package leetcode

//就是一个搜索问题啊
func searchInsert(nums []int, target int) int {
	
	for idx,v := range nums {
		if v==target{
			return idx
		}else if v<target{
			//如果这个小，而下一个大呢，如果最后一个都小呢
			if idx+1<len(nums) && nums[idx+1]>=target{
				return idx+1
			}else if idx==len(nums)-1{
				return idx + 1
			}
		}
	}
	return 0
}
