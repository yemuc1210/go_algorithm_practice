package offerS74

import "sort"

/*  leetcode 56  数组   排序
剑指 Offer II 074. 合并区间
以数组 intervals 表示若干个区间的集合，
其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
*/

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
		if pre[1] < cur[0] { //不重合
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