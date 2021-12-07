package lt287
/*中
287. 寻找重复数
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。

你设计的解决方案必须不修改数组 nums 且只用常量级 O(1) 的额外空间。
*/
/*
假设只有一个重复的整数
map 排序
求和 
快慢指针
*/
func findDuplicate(nums []int) int {
	res :=0
	for fast:=0;res!=fast || fast==0; {
		res = nums[res]
		fast = nums[nums[fast]]
	}
	for i:=0;res!=i;i=nums[i] {
		res = nums[res]
	}
	return res
}