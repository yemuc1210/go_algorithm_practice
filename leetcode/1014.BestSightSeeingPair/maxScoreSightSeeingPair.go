package lt1014
/*
给你一个正整数数组 values，其中 values[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的 距离 为 j - i。
一对景点（i < j）组成的观光组合的得分为 values[i] + values[j] + i - j ，也就是景点的评分之和 减去 它们两者之间的距离。
返回一对观光景点能取得的最高分。
输入：values = [8,1,5,2,6]
输出：11
解释：i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11
*/
func maxScoreSightseeingPair(values []int) int {
    ans, mx := 0, values[0] + 0
    for j := 1; j < len(values); j++ {
        ans = max(ans, mx + values[j] - j)
        // 边遍历边维护
        mx = max(mx, values[j] + j)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
