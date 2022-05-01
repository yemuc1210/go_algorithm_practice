package tencent50
/*简单
136. 只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
*/
/*
map
异或
*/
func singleNumber(nums []int) int {
	//题目假设数组非空
	res := nums[0]
	for i:=1;i<len(nums);i++{
		res ^= nums[i]
	}
	return res
}