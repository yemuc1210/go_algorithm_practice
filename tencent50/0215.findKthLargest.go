package tencent50

import "sort"

/*中  记得回顾
215. 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
*/
/*
思路1：排序

思路2：插入排序 从大到小排
*/
// 解法一 排序，排序的方法反而速度是最快的
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 解法二 这个方法的理论依据是 partition 得到的点的下标就是最终排序之后的下标，根据这个下标，我们可以判断第 K 大的数在哪里
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
			return 0
	}
	return selection(nums, 0, len(nums)-1, len(nums)-k)  //第k大，也就是N-k小
}

func selection(arr []int, l, r, k int) int {
	if l == r {
			return arr[l]
	}
	p := partition164(arr, l, r)
	if k == p {
			return arr[p]
	} else if k < p {
			return selection(arr, l, p-1, k)
	} else {
			return selection(arr, p+1, r, k)
	}
}

func partition164(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
			if a[j] < pivot {
					i++
					a[j], a[i] = a[i], a[j]
			}
	}
	a[i+1], a[hi] = a[hi], a[i+1]  //a[i]左边都是小于它的
	return i + 1
}



//基于堆排序的方案
func findKthLargest2(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)  //建立大根堆
	for i := len(nums) - 1; i >= len(nums) - k + 1; i-- {  //做k-1次删除操作后，堆顶元素就是
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize/2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i * 2 + 1, i * 2 + 2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}