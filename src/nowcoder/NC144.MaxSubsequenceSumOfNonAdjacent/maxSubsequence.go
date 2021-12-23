package NC144



/*
描述
给你一个数组，其长度为 n  ，在其中选出一个子序列，
子序列中任意两个数不能有相邻的下标（子序列可以为空）

本题中子序列指在数组中任意挑选若干个数组成的数组。

示例1
输入：
3,[1,2,3]
返回值：
4
说明：
有[],[1],[2],[3],[1,3] 4种选取方式其中[1,3]选取最优，答案为4
*/

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算  dp 打家劫舍问题
 * @param n int整型 数组的长度
 * @param array int整型一维数组 长度为n的数组
 * @return long长整型
 */

/*
设置一个状态转移数组dp，dp[i]表示数组中前i个元素所能偷的最大金额是多少

状态转移表达式：
(1)对于当前的元素arr[i],如果偷，那么dp[i] = dp[i-2] + arr[i]
(2)如果不偷，那么dp[i] = dp[i-1]
*/
func subsequence( n int ,  array []int ) int64 {
    // write code here
    if len(array)<2{
        return max(0,int64(array[0]))
    }
	dp := make([]int64, n)

	dp[0] = max(0,int64(array[0]))
	dp[1] = max(dp[0],int64(array[1]))
	for i:=2;i<n;i++{
		dp[i] = max(dp[i-1],dp[i-2]+int64(array[i]))
	}
	return dp[n-1]
}

func max(a,b int64)int64{
	if a>b {
		return a
	}
	return b
}