package offerS4
/*
剑指 Offer II 004. 只出现一次的数字 
给你一个整数数组 nums ，除某个元素仅出现 一次 外，
其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。


剑指offer 56-2

异或
*/
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}