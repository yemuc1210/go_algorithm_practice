package lt162
/*
峰值元素是指其值大于左右相邻值的元素。

给你一个输入数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞ 。


返回任意一个峰值即可

规律
规律一：如果nums[i] > nums[i+1]，则在i之前一定存在峰值元素

规律二：如果nums[i] < nums[i+1]，则在i+1之后一定存在峰值元素
*/
func findPeakElement(nums []int) int {

	left,right := 0,len(nums)-1

	for left<right{
		mid := left + (right - left)/2

		if nums[mid]>nums[mid+1]{
			//左边
			right = mid
		}else{
			left=mid+1
		}

	}
	return left
}