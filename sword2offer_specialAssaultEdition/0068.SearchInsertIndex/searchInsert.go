package offerS68
/*
剑指 Offer II 068. 查找插入位置
给定一个排序的整数数组 nums 和一个整数目标值 target ，
请在数组中找到 target ，并返回其下标。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

二分
*/
func searchInsert(nums []int, target int) int {
	if len(nums) == 0{
		return -1
	}
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