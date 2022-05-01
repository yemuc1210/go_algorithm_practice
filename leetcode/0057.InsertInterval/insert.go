package lt57

import "sort"

/* 中
57. 插入区间
给你一个 无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
*/
/*
leetcode 56合并区间

*/
func insert(intervals [][]int, newInterval []int) [][]int {
	//先添加
	intervals = append(intervals, newInterval)
	//然后使用lt56的逻辑合并
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	if len(intervals)==0 || len(intervals)==1{
		return intervals
	}
	//首先排序，根据左端点排序
	//使用sort.Slice接口，自定义排序方法
	sort.Slice(intervals, func(i,j int)bool{
		//根据自定义的第二个参数less函数 排序
		return intervals[i][0] < intervals[j][0]
	})  //根据左端点大小排序，小的在前
	
	res := [][]int{}

	//判断是否有重合，有则合并
	pre := intervals[0]
	for i:=1;i<len(intervals);i++{
		cur := intervals[i]
		//尝试合并
		if pre[1] < cur[0] { //不重合 前一个末端点 和 后一个左端点
			res = append(res, pre)
			pre = cur
		}else { //重合
			pre[1] = max(pre[1],cur[1])  //新的右端点
			//重合，合并，暂时不加入，只有不能合并的加入
		}
	}
	res = append(res, pre)  //最后一个必然加入
	return res
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}