package NC37

import (
	"sort"
	. "structs"
)

/*
描述
给出一组区间，请合并所有重叠的区间。
请保证合并后的区间按区间起点升序排列。

数据范围：区间组数 0 \le n \le 10000≤n≤1000，区间内 的值都满足 0 \le val \le 100000≤val≤10000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)
进阶：空间复杂度 O(val)O(val)，时间复杂度O(val)O(val)
*/
/*
 * type Interval struct {
 *   Start int
 *   End int
 * }
 */

/**
 * 合并区间 lt56
 * @param intervals Interval类一维数组
 * @return Interval类一维数组
 */
func merge( intervals []*Interval ) []*Interval {
	if len(intervals)==0 || len(intervals)==1{
		return intervals
	}
	// 首先进行排序 根据左端点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	res := []*Interval{}
	//合并 
	pre,cur := intervals[0], &Interval{}
	for i:=1;i<len(intervals);i++{
		cur = intervals[i]
		if pre.End < cur.Start {
			//不重合
			res = append(res, pre)
			pre = cur
		}else{
			//重合
			pre.End = max(pre.End, cur.End)
		}
	}
	//最后一个
	res = append(res, pre)
	return res
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}