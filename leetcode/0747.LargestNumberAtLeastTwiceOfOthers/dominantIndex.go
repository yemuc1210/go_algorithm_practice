package lt747

import (
	"fmt"
	"math"
)

/*
747. 至少是其他数字两倍的最大数

给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。

请你找出数组中的最大元素并检查它是否 至少 是数组中每个其他数字的两倍 。
如果是，则返回 最大元素的下标 ，否则返回 -1 。
*/
//找出最大的和第二大的
func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	type pair struct{num,idx int}
	maxNum := pair{nums[0],0}
	
	//分为两次遍历
	for i,v := range nums {
		if v>maxNum.num {
			maxNum.num = v
			maxNum.idx = i
		}
	}
	//找出第二大
	secMaxNum := pair{math.MinInt32,-1}
	for i,v := range nums {
		if i!=maxNum.idx{
			if v>secMaxNum.num && v < maxNum.num {
				secMaxNum.idx = i
				secMaxNum.num = v
			}
		}
	}
	fmt.Println(maxNum,secMaxNum)
	//secondMax.num可能是0
	if secMaxNum.num == 0 {
		return maxNum.idx
	}
	if maxNum.num/secMaxNum.num >= 2 {
		return maxNum.idx
	}
	return -1
}

//一次遍历解决
func dominantIndex2(nums []int) int {
	max,idx,less := 0,0,1
	for i:=0;i<len(nums);i++{
		if nums[i]>max {
			less = max  //记录上一个max
			max = nums[i]
			idx = i
		}else if nums[i] > less {
			// nums[i] <=max && nums[i] >less  更新第二大
			less = nums[i]
		}
	}
	if max>=less*2 { 
		return idx
	}
	return -1
}