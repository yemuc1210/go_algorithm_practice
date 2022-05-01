package main

import "math/rand"

type Solution1 []int

func Constructo1r(nums []int) Solution1 {
    return nums
}

func (nums Solution1) Pick1(target int) (ans int) {
    cnt := 0
    for i, num := range nums {
        if num == target {
            cnt++ // 第 cnt 次遇到 target
            if rand.Intn(cnt) == 0 {
                ans = i
            }
        }
    }
    return
}

