package lt453
/*
453. 最小操作次数使数组元素相等
给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。
返回让数组所有元素相等的最小操作次数。

 
n-1个数同时加一，就好比每次有一个数自身减一，因为只能做减法，
所以数组最后的数只能是最小值。这样的话每个元素减去最小值求其和就是答案

数学题

假设目前数组总和为sum，我们需要移动次数为m，那么整体数组总和将会增加m * (n - 1)，这里的n为数组长度，最后数组所有元素都相等为x，于是有：

sum + m * (n - 1) = x * n (1)

我们再设数组最小的元素为min_val，m = x - min_val​，即 ​x = m + min_val​带入(1)得：

m = sum - min_val * n​
*/
func minMoves(nums []int) int {
	sum := 0
	minmum := func(nums []int) int {
		min := nums[0]
		for i:=1;i<len(nums);i++{
			if nums[i]<min{
				min = nums[i]
			}
		}
		return min
	}(nums)
	for _,v := range nums{
		sum += v-minmum
	}
	return sum
}
