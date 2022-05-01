package lt594

/*简单
594. 最长和谐子序列
和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。

现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。

数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。
*/
/*
1. 用二维数组记录当前元素与其他元素之间的差值
	形成子序列的过程，就沿着右下方向蔓延
	dfs搜索
3. dp  搞错了  【要求是最大值最小值差值1】
	dp[i]：[0:i]的最长子序列长度
		需要last记录子序列最后一个元素
	状态转移方程：dp[i] = dp[i-1] + 1   if nums[i] - last <= 1
		else dp[i] = dp[i-1]
3. map  统计数字频次 然后找差值1的两个数组的频次和
动态的维护两个数的频次和就是最后要求的子数组的最大长度
*/
func findLHS(nums []int) int {
	if len(nums) < 2 { //0 1
		return 0
	}
	m := make(map[int]int,len(nums))
	for _,num := range nums{
		//统计频次
		if _,exist := m[num];exist {
			m[num]++
			continue
		}
		m[num] = 1
	}

	res := 0 // 最大长度
	for k,cnt := range m {
		//找差值为1 的数组
		if n,exist := m[k+1];exist {
			if cnt + n > res {
				//动态更新长度
				res = cnt + n
			}
		}
	}
	return res
}