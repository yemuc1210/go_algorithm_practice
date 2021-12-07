package offer56
/*
在一个数组 nums 中除一个数字只出现一次之外，
其他数字都出现了三次。请找出那个只出现一次的数字。

示例 1：
输入：nums = [3,4,3,3]
输出：4

时间复杂度O(n)，则不能使用排序
空间复杂度O(1)，不能使用map


一个数和自身异或，等于自身
*/
func singleNumbers(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}