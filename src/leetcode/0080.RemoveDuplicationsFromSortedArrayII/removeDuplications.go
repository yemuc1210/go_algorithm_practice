package lt80
//lt26  
/*
80. 删除有序数组中的重复项 II
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/

/*
原地删除  允许重复，最多出现两次
*/
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	idx := 0  //idx遍历原数组
	for _,v := range nums{
		if idx<2 || v > nums[idx-2] { 
			nums[idx] = v
			idx++
		}
	}
	return idx
}
