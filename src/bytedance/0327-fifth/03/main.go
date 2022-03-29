package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var x int
	var n int
	fmt.Scanln(&x, &n) // 初始x ，经过时间n

	h := &IHeap{}
	heap.Init(h)
	// 第0秒也会操作
	if x%2 == 0 {
		// 分成两半
		h.Push(x / 2)
		h.Push(x / 2)
	} else {
		h.Push(x)
	}
	for i := 1; i <= n; i++ {
		// 寻找当前长度里面为偶数的最长的
		fmt.Println(h,h.Len())
		h1 := &IHeap{}
		heap.Init(h1)
		for j := 0; j < h.Len(); j++ {
			x := h.Pop().(int)
			h1.Push(x + 1)
		}
		fmt.Println(h,h1)
		// 进行切断
		var f bool
		for j := 0; j < h1.Len(); j++ {
			x := h1.Pop().(int)
			if x%2 == 0 && !f {
				h.Push(x / 2)
				h.Push(x / 2)
				f = true
			} else {
				h.Push(x)
			}
		}
		fmt.Println(h,h.Len(),h1,h1.Len())
	}
	fmt.Println(h.Pop())
}

/*
1 6
1 
1 1
1 1 2
1 1 2 3

*/

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
