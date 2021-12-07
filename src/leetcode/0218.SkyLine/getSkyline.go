package lt218

import "sort"

/*
城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。
给你所有建筑物的位置和高度，请返回由这些建筑物形成的 天际线 。

每个建筑物的几何信息由数组 buildings 表示，
其中三元组 buildings[i] = [lefti, righti, heighti] 表示：

lefti 是第 i 座建筑物左边缘的 x 坐标。
righti 是第 i 座建筑物右边缘的 x 坐标。
heighti 是第 i 座建筑物的高度。
天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，
并按 x 坐标 进行 排序 。
关键点是水平线段的左端点。
列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，
仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

注意：输出天际线中不得有连续的相同高度的水平线。
例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；
三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]

*/
func getSkyline(buildings [][]int) [][]int {
	size := len(buildings)
	es := make([]E, 0)
	for i, b := range buildings {
		l := b[0]
		r := b[1]
		h := b[2]
		// 1-- enter
		el := NewE(i, l, h, 0)
		es = append(es, el)
		// 0 -- leave
		er := NewE(i, r, h, 1)
		es = append(es, er)
	}
	skyline := make([][]int, 0)
	sort.Slice(es, func(i, j int) bool {
		if es[i].X == es[j].X {
			if es[i].T == es[j].T {
				if es[i].T == 0 {
					return es[i].H > es[j].H
				}
				return es[i].H < es[j].H
			}
			return es[i].T < es[j].T
		}
		return es[i].X < es[j].X
	})
	pq := NewIndexMaxPQ(size)
	for _, e := range es {
		curH := pq.Front()
		if e.T == 0 {
			if e.H > curH {
				skyline = append(skyline, []int{e.X, e.H})
			}
			pq.Enque(e.N, e.H)
		} else {
			pq.Remove(e.N)
			h := pq.Front()
			if curH > h {
				skyline = append(skyline, []int{e.X, h})
			}
		}
	}
	return skyline
}

// E define
type E struct { // 定义一个 event 事件
	N int // number 编号
	X int // x 坐标
	H int // height 高度
	T int // type  0-进入 1-离开
}

// NewE define
func NewE(n, x, h, t int) E {
	return E{
		N: n,
		X: x,
		H: h,
		T: t,
	}
}
// IndexMaxPQ define
type IndexMaxPQ struct {
	items []int
	pq    []int
	qp    []int
	total int
}
// NewIndexMaxPQ define
func NewIndexMaxPQ(n int) IndexMaxPQ {
	qp := make([]int, n)
	for i := 0; i < n; i++ {
		qp[i] = -1
	}
	return IndexMaxPQ{
		items: make([]int, n),
		pq:    make([]int, n+1),
		qp:    qp,
	}
}

// Enque define
func (q *IndexMaxPQ) Enque(key, val int) {
	q.total++
	q.items[key] = val
	q.pq[q.total] = key
	q.qp[key] = q.total
	q.swim(q.total)
}

// Front define
func (q *IndexMaxPQ) Front() int {
	if q.total < 1 {
		return 0
	}
	return q.items[q.pq[1]]
}

// Remove define
func (q *IndexMaxPQ) Remove(key int) {
	rank := q.qp[key]
	q.exch(rank, q.total)
	q.total--
	q.qp[key] = -1
	q.sink(rank)
}

func (q *IndexMaxPQ) sink(n int) {
	for 2*n <= q.total {
		k := 2 * n
		if k < q.total && q.less(k, k+1) {
			k++
		}
		if q.less(k, n) {
			break
		}
		q.exch(k, n)
		n = k
	}
}

func (q *IndexMaxPQ) swim(n int) {
	for n > 1 {
		k := n / 2
		if q.less(n, k) {
			break
		}
		q.exch(n, k)
		n = k
	}
}

func (q *IndexMaxPQ) exch(i, j int) {
	q.pq[i], q.pq[j] = q.pq[j], q.pq[i]
	q.qp[q.pq[i]] = i
	q.qp[q.pq[j]] = j
}

func (q *IndexMaxPQ) less(i, j int) bool {
	return q.items[q.pq[i]] < q.items[q.pq[j]]
}
