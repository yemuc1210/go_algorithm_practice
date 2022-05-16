package main

import (
	"fmt"
	"sort"
)

func largestCombination(candidates []int) int {
	if len(candidates) == 1 && candidates[0] > 0 {
		return 1
	}
	// 组合
	sort.Slice(candidates,func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	n := len(candidates)
	res := 0
	totalAnSum := candidates[0]
	// 从前往后
	prefixSum := make([]int,n)
	for i:=1;i<n;i++ {
		prefixSum[i] = totalAnSum   // 不包含自身的前缀和
		totalAnSum &= candidates[i]
	}
	if totalAnSum > 0 {
		return n
	}
	// 从后往前
	tSum := candidates[n-1]
	postfixSum := make([]int,n)
	curLen := 0
	for i:=n-2;i>=0;i--{
		postfixSum[i] = tSum   // 不包含自身的后缀和
		tSum &= candidates[i]
	}
	// 否则 前缀和的思路
	curSum := candidates[0]
	for i:=1;i<n;i++{
		if curSum == 0 {
			if curLen>res {
				res = curLen
			}
			curSum = 1
			curLen = 0
			continue
		}
		fmt.Println(curSum,candidates[i],curSum&candidates[i])
		if curSum & candidates[i] > 0 {
			curSum &= candidates[i]
			curLen ++
			continue
		} 
		// 否则
		// 判断剩余的数字是否可能令结果为正
		if postfixSum[i] & curSum < 0 {
			if curLen>res {
				res = curLen
			}
			curSum = 1
			curLen = 0
			continue
		}
	}
	return res
}


func solve(nums []int) int {
	if len(nums) == 1 && nums[0] > 0 {
		return 1
	}
	sort.Slice(nums,func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res int
	// 搜索
	n := len(nums)
	var dfs func(i int,curSum int,curPathLen int)
	dfs = func(i int, curSum int,curPathLen int) {
		if i==n {
			if curPathLen > res {
				res = curPathLen
			}
			return 
		}
		// 访问
		tmpSum := curSum & nums[i]
		if tmpSum > 0 {
			dfs(i+1, tmpSum,curPathLen+1)
		}else {
			dfs(i+1, curSum,curPathLen)
		}
	}
	dfs(0,1,0)
	return res
}
func main() {
	candidates := []int{16,17,71,62,12,24,14}
	res := largestCombination(candidates)
	fmt.Println(res)
	res1 := solve(candidates)
	fmt.Println(res1)
}