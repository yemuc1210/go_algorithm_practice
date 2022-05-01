package tencent50
/*
238. 除自身以外数组的乘积
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，
其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。

说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

*/

/*
先计算累乘，然后➗当前元素  pass
计算每个数左右的累乘
*/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	left, right := 1, 1
	res := []int{}
	for i := 0; i < n; i++ {
		res = append(res, 1)
	}
	for i := 0; i < n; i++ {
		res[i] *= left
		left *= nums[i]

		res[n-i-1] *= right
		right *= nums[n-i-1]
	}
	return res
}