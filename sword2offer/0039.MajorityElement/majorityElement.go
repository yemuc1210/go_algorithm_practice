package offer39

/*
剑指 Offer 39. 数组中出现次数超过一半的数字
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

*/
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	curMarjority, curCnt := nums[0],1

	for i:=1;i<len(nums);i++{
		if nums[i] == curMarjority{
			curCnt ++
		}else {
			curCnt --
		}

		if curCnt == 0{
			curMarjority = nums[i]
			curCnt = 1
		}

	}
	return curMarjority
}