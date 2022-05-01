package lt502

import (
	"container/heap"
	"sort"
)

/*
502. IPO
假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，
力扣 希望在 IPO 之前开展一些项目以增加其资本。
由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。
帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。

给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。

最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。

总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。

答案保证在 32 位有符号整数范围内。


直觉是贪心法
根据资本排序，手中资本会增加profit[i]
	如果初始资本w>=max(captial)，则直接返回利润中最大的k的元素，因为有钱任性
	需要贪心地保证每次获取的利润最大
	所需资本从小到大排序，每次选择：当前持有w，从资本投入小于等于w的那部分中选择利润最大的项目j。
		更新手中资本w+profit[j]
		迭代w = w+profit[j]

数据结构：大根堆，将利润压入堆在，每次去做最大值



*/
func findMaximizedCapital(k, w int, profits, capital []int) int {
    n := len(profits)
    type pair struct{ c, p int }
    arr := make([]pair, n)
    for i, p := range profits {
        arr[i] = pair{capital[i], p}
    }
    sort.Slice(arr, func(i, j int) bool { return arr[i].c < arr[j].c })

    h := &hp{}
    for cur := 0; k > 0; k-- {
        for cur < n && arr[cur].c <= w {
            heap.Push(h, arr[cur].p)
            cur++
        }
        if h.Len() == 0 {
            break
        }
        w += heap.Pop(h).(int)
    }
    return w
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }   //从小到大排序
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
