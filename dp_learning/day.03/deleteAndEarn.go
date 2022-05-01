package dsday03
/*
给你一个整数数组 nums ，你可以对它进行一些操作。

每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，
你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。

开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。

*/
func deleteAndEarn(nums []int) int {
    maxVal := 0
    for _, val := range nums {
        maxVal = max(maxVal, val)
    }
    sum := make([]int, maxVal+1)
    for _, val := range nums {
        sum[val] += val
    }
    return rob(sum)
}

func rob(nums []int) int {
    first, second := nums[0], max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        first, second = second, max(first+nums[i], second)
    }
    return second
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

