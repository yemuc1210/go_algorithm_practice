package tencent50
/*简单
169. 多数元素
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
/*
136  只出现一个数字
由于假定数组非空，总是存在多数，并且只用返回一个
思路1：排序，输出中间元素
	手写排序
思路2：map计算
思路3：摩尔投票	
	同则+，异则-
*/
//摩尔投票
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}  //不过假设是非空
	votes := 1
	curElem := nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i] == curElem {
			votes ++
		}else {
			votes -- 
			if votes == 0 {
				curElem = nums[i]
				votes = 1
			}
		}
	}
	return curElem
}