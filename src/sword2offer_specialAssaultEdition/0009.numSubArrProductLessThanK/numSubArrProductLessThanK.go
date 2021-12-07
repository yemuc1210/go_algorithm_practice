package offerS9
/*
剑指 Offer II 009. 乘积小于 K 的子数组
给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。

连续子数组  应该是滑动窗口
窗口内每加入一个数，如果乘积仍然小于k，则 满足条件的子数组数量 += 窗口内元素数量
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	if n==0{
		return 0
	}
	res := 0

	left,right := 0,0
	curProduct := 1
	for right < n{
		num := nums[right]
		// num := nums[right]
		right ++
		//当前元素大于k，从右边重新开始
		if num >= k{
			left = right
			curProduct = 1
			continue
		}
		curProduct = curProduct * num
		//乘积大于等于k,左边缩小
		for curProduct >= k {
			curProduct = curProduct / nums[left]
			left++
		}
		if curProduct < k{
			res += (right-left)
		}
	}
	return res
}