package lt413

/*
数组 A 包含 N 个数，且索引从0开始。数组 A 的一个子数组划分为数组 (P, Q)，P 与 Q 是整数且满足 0<=P<Q<N 。

如果满足以下条件，则称子数组(P, Q)为等差数组：

元素 A[P], A[p + 1], ..., A[Q - 1], A[Q] 是等差的。并且 P + 1 < Q 。

函数要返回数组 A 中所有为等差数组的子数组个数。

每判断一组 3 个连续的数列，只需要用一个变量累加前面已经有多少个满足题意的连续元素，
只要满足题意的等差数列就加上这个累加值。一旦不满足等差的条件，累加值置 0。如此循环一次即可找到题目要求的答案。
*/
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	res, dp := 0, 0
	for i := 1; i < len(nums)-1; i++ {  // 0<=P<Q<N
		if nums[i+1]-nums[i] == nums[i]-nums[i-1] {   //返回子数组个数
			dp++
			res += dp
		} else {
			dp = 0
		}
	}
	return res
}
