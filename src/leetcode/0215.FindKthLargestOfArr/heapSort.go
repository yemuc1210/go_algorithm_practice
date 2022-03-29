package lt215

import "container/heap"

// 手写堆排序
func findKthLargest215(nums []int, k int) int {
	heapSize := len(nums)
	// 建堆
	buildMaxHeap1(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify1(nums, 0, heapSize)
	}
	return nums[0]
}
func buildMaxHeap1(a []int, heapSize int) {
	// 从最后一个数字开始调整
	for i := heapSize / 2; i >= 0; i-- {
		// 堆调整
		maxHeapify1(a, i, heapSize)
	}
}

func maxHeapify1(a []int, i, heapSize int) {
	// 左右结点
	l, r := 2*i+1, 2*i+2
	largerOne := i
	if l < heapSize && a[l] > a[largerOne] {
		largerOne = l
	}
	if r < heapSize && a[r] > a[largerOne] {
		largerOne = r
	}
	if largerOne != i {
		//进行调整
		a[i], a[largerOne] = a[largerOne], a[i]
		// 递归向下调整
		maxHeapify1(a, largerOne, heapSize)
	}
	// 否则，以当前结点i为根的符合大根堆的定义
}

// 依赖于go提供的接口
func findK(nums []int, k int) int {
	h := &Iheap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}
	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}

type Iheap []int

func (h Iheap) Len() int { return len(h) }

// 大根堆需要变
func (h Iheap) Less(i, j int) bool { return h[i] > h[j] }
func (h Iheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Iheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Iheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
