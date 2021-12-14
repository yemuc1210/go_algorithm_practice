package lt630

import (
	"container/heap"
	"sort"
)

/* lt207课程表    lt210课程表II
630. 课程表 III
这里有 n 门不同的在线课程，按从 1 到 n 编号。
给你一个数组 courses ，其中 courses[i] = [durationi, lastDayi]
	表示第 i 门课将会 持续 上 durationi 天课，
并且必须在不晚于 lastDayi 的时候完成。

你的学期从第 1 天开始。且不能同时修读两门及两门以上的课程。

返回你最多可以修读的课程数目。
*/
/*
这不就是aoe嘛
不能同时修读两门或两门以上
方法：优先队列+贪心
规律1 ：有两门（t1,d1)(t2,d2)若d1<=d2，则最优做法是先学习t1，然后t2
所以，先按关闭时间d升序
*/
func scheduleCourse(courses [][]int) int {
    sort.Slice(courses, func(i, j int) bool {
        return courses[i][1] < courses[j][1]
    })

    h := &Heap{}
    total := 0 // 优先队列中所有课程的总时间
    for _, course := range courses {
        if t := course[0]; total+t <= course[1] {
            total += t
            heap.Push(h, t)
        } else if h.Len() > 0 && t < h.IntSlice[0] {
            total += t - h.IntSlice[0]
            h.IntSlice[0] = t
            heap.Fix(h, 0)
        }
    }
    return h.Len()
}

type Heap struct {
    sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
    return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
    h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
    a := h.IntSlice
    x := a[len(a)-1]
    h.IntSlice = a[:len(a)-1]
    return x
}
