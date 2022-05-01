package oofer66

//LeetCode  0238

func constructArr(nums []int) []int {
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