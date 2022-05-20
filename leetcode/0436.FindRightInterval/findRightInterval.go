package main

import "sort"

// intervals[i] = [start_i, end_i] 每个start不同
// 区间i的右侧区块为j，有start_j >= end_i，且start_j最小化
// 也就是最近的右区间2
func findRightInterval(intervals [][]int) []int {
	if len(intervals) == 1 {
		return []int{-1}
	}
	for i:=range intervals {
		intervals[i] = append(intervals[i], i)  //将下标加入
	}
	sort.Slice(intervals,func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]   // 按照start排序
	})
	n := len(intervals)
	res := make([]int, len(intervals))
	// 对于每个区间，二分查找
	for _,interval := range intervals {
		i := sort.Search(n, func(i int) bool {
			return intervals[i][0] >= interval[1] // start_j >= end_i
		})
		if i<n {
			res[interval[2]] = intervals[i][2]
		}else{
			res[interval[2]] = -1
		}
	}
	return res

}