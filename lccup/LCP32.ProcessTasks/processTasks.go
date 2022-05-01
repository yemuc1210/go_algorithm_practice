package lcp32
/*
LCP 32. 批量处理任务
某实验室计算机待处理任务以 [start,end,period] 格式记于二维数组 tasks，
表示完成该任务的时间范围为起始时间 start 至结束时间 end 之间，需要计算机投入 period 的时长，
注意：

period 可为不连续时间
首尾时间均包含在内
处于开机状态的计算机可同时处理任意多个任务，请返回电脑最少开机多久，可处理完所有任务。

贪心法  任务看作小段



*/
import (
	"container/heap"
	"sort"
)

type Pair struct {
	a, b, index int
}

type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].a < pq[j].a
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Pair)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func processTasks(tasks [][]int) int {
	var vt []int

	n := len(tasks)
	ts := make(map[int]bool)
	for i := 0; i < n; i++ {
		ts[tasks[i][0]] = true
		ts[tasks[i][1] + 1] = true
	}
	m := len(ts)
	for key := range ts {
		vt = append(vt, key)
	}
	sort.Ints(vt)
	mp := make(map[int]int)
	for i := 0; i < m; i++ {
		mp[vt[i]] = i
	}

	starts := make([][]int, m)
	for i := 0; i < n; i++ {
		starts[mp[tasks[i][0]]] = append(starts[mp[tasks[i][0]]], i)
	}

	ans := 0
	extra := 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for i := 0; i < m - 1; i++ {
		for len(pq) > 0 {
			if tasks[pq[0].b][1] < vt[i] {
				heap.Pop(&pq)
			} else {
				break
			}
		}

		for _, u := range starts[i] {
			heap.Push(&pq, &Pair{a: extra + tasks[u][1] - vt[i] + 1 - tasks[u][2], b: u})
		}

		currentSeg := vt[i + 1] - vt[i]
		if len(pq) > 0 {
			minSlots := pq[0].a - extra
			slotsToDel := currentSeg
			if minSlots < currentSeg {
				delta := currentSeg - minSlots
				ans += delta
				slotsToDel -= delta
			}
			extra += slotsToDel
		}
	}

	return ans
}
