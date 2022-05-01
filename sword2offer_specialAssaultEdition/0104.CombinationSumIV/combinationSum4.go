package offerS104
/* lt377  dp
剑指 Offer II 104. 排列的数目
给定一个由 不同 正整数组成的数组 nums ，和一个目标整数 target 。
请从 nums 中找出并返回总和为 target 的元素组合的个数。
数组中的数字可以在一次排列中出现任意次，但是顺序不同的序列被视作不同的组合。

题目数据保证答案符合 32 位整数范围。
*/

/*
和上一题只求最小数不同，这里需要得到组合个数
dp[i]  得到i的组合数
dp[i] = dp[i-num] + dp[i]   //不取num  和 取num
*/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i:=1;i<=target;i++{
		for _,num := range nums {
			if num <= i {
				dp[i] = dp[i] + dp[i-num]
			}
		}
	}
	return dp[target]
}