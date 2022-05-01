package leetcode

import "sort"

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
。
*/


func merge(intervals [][]int) [][]int {
	if len(intervals)==0 || len(intervals)==1{
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})  //使用内置的sort.slice接口实现数组的排序
	res :=[][]int{}
	pre:= intervals[0]

	for i:=1;i<len(intervals);i++{
		cur := intervals[i]
		//尝试cur能不能和pre合并
		if pre[1] < cur[0] {//不重合
			res = append(res, pre)
			pre = cur
		}else {//存在重合
			pre[1] = max(pre[1],cur[1])
		}
	}
	res = append(res, pre)
	return res
}
func max(a, b int) int {
	if a > b { return a }
	return b
}