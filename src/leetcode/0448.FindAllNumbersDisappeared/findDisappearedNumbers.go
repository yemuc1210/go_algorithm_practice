package lt448
/*
448. 找到所有数组中消失的数字
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。
请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，
并以数组的形式返回结果。
*/
/*
要求不能使用额外的空间，那么只能想办法在原有数组上进行修改，并且这个修改是可还原的。时间复杂度也只能允许我们一层循环。只要循环一次能标记出已经出现过的数字，这道题就可以按要求解答出来。这里笔者的标记方法是把 |nums[i]|-1 索引位置的元素标记为负数。即 nums[| nums[i] |- 1] * -1。这里需要注意的是，nums[i] 需要加绝对值，因为它可能被之前的数置为负数了，需要还原一下。最后再遍历一次数组，若当前数组元素 nums[i] 为负数，说明我们在数组中存在数字 i+1。把结果输出到最终数组里即可。
*/
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}