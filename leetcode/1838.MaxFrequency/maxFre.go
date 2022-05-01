package lt1838

import "sort"

/*
元素的 频数 是该元素在一个数组中出现的次数。

给你一个整数数组 nums 和一个整数 k 。
在一步操作中，你可以选择 nums 的一个下标，
并将该下标对应元素的值增加 1 。

执行最多 k 次操作后，返回数组中最高频元素的 最大可能频数 。

*/

//找到中位数，最大频次数
//排序+滑动窗口
func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    ans := 1
    for l, r, total := 0, 1, 0; r < len(nums); r++ {
        total += (nums[r] - nums[r-1]) * (r - l)
        for total > k {
            total -= nums[r] - nums[l]
            l++
        }
        ans = max(ans, r-l+1)
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}