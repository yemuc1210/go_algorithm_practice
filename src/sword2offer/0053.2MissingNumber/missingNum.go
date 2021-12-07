package offer53
/*
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
*/
func missingNumber(nums []int) int {
	//递增排序 唯一  长度0~n-1,数字范围0~n-1
	n := len(nums)
	for i:=0;i<n;i++{
		if nums[i]!=i{
			return i
		}
	}
	return n
}