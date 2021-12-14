package lt136

/*
给定一个非空整数数组，除了某个元素只出现一次以外，
其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

排序，然后遍历    非线性时间复杂度

map存储计数   

异或运算的性质：任何一个数字异或它自己都等于0。
也就是说，如果我们从头到尾依次异或数组中的每一个数字，
那么最终的结果刚好是那个只出现依次的数字，
因为那些出现两次的数字全部在异或中抵消掉了。
于是最终做法是从头到尾依次异或数组中的每一个数字，
那么最终得到的结果就是两个只出现一次的数字的异或结果。
因为其他数字都出现了两次，在异或中全部抵消掉了。**利用的性质是 x^x = 0**。
*/

func singleNumber(nums []int) int {
	//已知数组非空，所以不用判空
	x := nums[0]
	for i:=1;i<len(nums);i++{
		x = x^nums[i]   //异或
	}
	return x
}