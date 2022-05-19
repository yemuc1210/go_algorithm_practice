package main

import "sort"

// 返回使得所有元素相等的最小移动数
// 中位数
func minMoves2(nums []int) (ans int) {
    sort.Ints(nums)
    x := nums[len(nums)/2]
    for _, num := range nums {
        ans += abs(num - x)
    }
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
