package main

import (
	"fmt"
	"sort"
)

func flipgame(fronts []int, backs []int) int {
	// 1. 正反相同的不行
	n := len(fronts)
	nums := make(map[int]struct{})
	for i:=0;i<n;i++{
		if fronts[i]==backs[i] {
			nums[fronts[i]] = struct{}{}
		}
	}
	nums1 := []int{}
	for i:=0;i<n;i++{
		if _,ok:=nums[fronts[i]];!ok {
			nums1 = append(nums1, fronts[i])
		}
		if _,ok:=nums[backs[i]];!ok {
			nums1 = append(nums1, backs[i])
		}
	}
	// 排序
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i]<nums1[j]
	})
	if len(nums1)>0{
		return nums1[0]
	}
	return 0
}

func main(){
	// fronts :=[]int{1,2,4,4,7}
	// backs := []int{1,3,4,1,3}
	fronts :=[]int{1}
	backs := []int{1}
	res := flipgame(fronts,backs)
	fmt.Println(res)
}