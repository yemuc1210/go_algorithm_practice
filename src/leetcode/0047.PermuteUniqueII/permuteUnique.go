package lt47

import "sort"
/*
算法计划  递归回溯

*/

func permuteUnique(nums []int) (ans [][]int) {
    sort.Ints(nums)
    n := len(nums)
    perm := []int{}
    vis := make([]bool, n)
    var backtrack func(int)
    backtrack = func(idx int) {
        if idx == n {
            ans = append(ans, append([]int(nil), perm...))
            return
        }
        for i, v := range nums {
            if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
                continue
            }
            perm = append(perm, v)
            vis[i] = true
            backtrack(idx + 1)
            vis[i] = false
            perm = perm[:len(perm)-1]
        }
    }
    backtrack(0)
    return
}