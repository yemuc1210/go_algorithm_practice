package lt918

import "math"

/*
给定一个由整数数组 A 表示的环形数组 C，
求 C 的非空子数组的最大可能和。

在此处，环形数组意味着数组的末端将会与开头相连呈环状。（
形式上，当0 <= i < A.length 时 C[i] = A[i]，且当 i >= 0 时 C[i+A.length] = C[i]）

此外，子数组最多只能包含固定缓冲区 A 中的每个元素一次。
（形式上，对于子数组 C[i], C[i+1], ..., C[j]，
不存在 i <= k1, k2 <= j 其中 k1 % A.length = k2 % A.length）
*/
func maxSubarraySumCircular(nums []int) int {
	if len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	n, sum := len(nums), 0
	for _, v := range nums {
		sum += v
	}
	kad := kadane(nums)
	for i := 0; i < n; i++ {
		nums[i] = -nums[i]
	}
	negativeMax := kadane(nums)
	if sum+negativeMax <= 0 {
		return kad
	}
	return max(kad, sum+negativeMax)
}

func kadane(a []int) int {
	n, MaxEndingHere, maxSoFar := len(a), a[0], math.MinInt32
	for i := 1; i < n; i++ {
		MaxEndingHere = max(a[i], MaxEndingHere+a[i])
		maxSoFar = max(MaxEndingHere, maxSoFar)
	}
	return maxSoFar
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
