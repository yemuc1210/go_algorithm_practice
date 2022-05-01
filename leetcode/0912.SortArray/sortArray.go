package main

import (
	"fmt"
)

// 手写快排进行排序
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left <= right {
		pivot := partition(nums, left, right)
		// 递归
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
}
func partition(nums []int, left, right int) int {
	pivot := left + (right-left)/2
	// 根据pivot进行划分
	i, j := left, right //用闭区间
	pivotElem := nums[pivot]
	for i <= j { // 闭区间，因此i==j是有意义的
		for i <= j && nums[i] <= pivotElem {
			i++
		}
		for i <= j && nums[j] > pivotElem {
			j--
		}
		// 那么，可以让pivot左边都是小于等于pivotElem的
		if i > j {
			break
		}
		// swap
		nums[i], nums[j] = nums[j], nums[i]
	}
	// pivotElem归位
	nums[j] = pivotElem // i>j  i指向是小于等于  j是大于
	return pivot
}

func main() {
	fmt.Println("vim-go")
	nums := []int{5, 1, 1, 2, 0, 0}
	nums = sortArray(nums)
	fmt.Println(nums)
}
