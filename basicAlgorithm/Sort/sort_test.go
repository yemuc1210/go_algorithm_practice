package sDemoSort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{3, 4, 2, 1, 2, 1}
	BubbleSort(nums) // 传入数组，传的是引用
	fmt.Println(nums)
}

func TestInsertSort(t *testing.T) {
	var nums []int = []int{3, 4, 2, 1, 2, 1}
	InsertSort(nums)
	fmt.Println(nums)
}

func TestQuickSort(t *testing.T) {
	var nums []int = []int{3, 4, 2, 1, 2, 1}
	QuickSort(nums)
	fmt.Println(nums)
}

func TestMergeSort(t *testing.T) {
	var nums []int = []int{3, 4, 2, 1, 2, 1}
	MergeSort(nums)
	fmt.Println(nums)
}

func TestHeapSort(t *testing.T) {
	var nums []int = []int{3, 4, 2, 1, 2, 1}
	HeapSort(nums)
	fmt.Println(nums)
}

func TestSelectSort(t *testing.T) {
	var nums []int = []int{3, 4, 2, 1, 2, 1}
	SelectSort(nums)
	fmt.Println(nums)
}

func TestBucketSort(t *testing.T) {
	BucketSort([]int{})
}
