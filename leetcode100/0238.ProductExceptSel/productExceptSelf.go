package lt238

/*
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，
其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

要求时间复杂度O(n),不使用除法
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
