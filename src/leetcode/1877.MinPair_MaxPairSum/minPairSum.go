package lt1877

import "sort"

/*
数组中最大数对和的最小值


给定长度为偶数n的数组，将其中元素分为n/2个数对，返回最小的 最大数对和

最大数对和：所有数对中，最大的

读起来真费劲，直接排序，首尾相加
*/
func minPairSum(nums []int) int {
	sort.Ints(nums)
	ans :=0
	n := len(nums)
    for i, val := range nums[:n/2] {
        ans = max(ans, val+nums[n-1-i])
    }
    return ans
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}