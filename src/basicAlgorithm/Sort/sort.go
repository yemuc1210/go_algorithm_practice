package sDemoSort

import (
	"container/heap"
	"math"
	"strconv"
)

// 实现一个sort包

// 定义数据类型
type DateType int

// 1.冒泡排序 稳定
func BubbleSort(nums []int) {
	// 每次遍历 大的放到最后
	// 双层循环
	for i := len(nums) - 1; i >= 1; i-- { // 每次遍历 可以确定i位置的值
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] { //max: j=i-1 j+1=i
				nums[j], nums[j+1] = nums[j+1], nums[j] //swap
			}
		}
	}

}

// 2.插入排序 不稳定的排序
func InsertSort(nums []int) {
	if len(nums) <= 1 {
		return //默认有序
	}
	// 默认第一个有序
	lastIdx := 0 // 有序部分的最后一个位置  可以省略
	for i := 1; i < len(nums); i++ {
		// 当前元素插入有序
		// 1. 搜索插入位置  使用二分法
		insertIdx := searchIdx(nums[:lastIdx+1], nums[i])
		// 2. 在insertIdx之后插入
		a, b, c := nums[:insertIdx+1], nums[insertIdx+1:i+1], nums[i+1:]
		// 对 b 操作
		b = b[:len(b)-1]                 //删除nums[i]
		b = append([]int{nums[i]}, b...) // 头插
		// 3. 组合
		nums = append(a, append(b, c...)...)
	}
}

// 搜索元素下标，若没有，则返回插入下标
func searchIdx(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 闭区间
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target { // 这样写，是不稳定的，因为可能多个相同的数在一起
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 如果没有返回，说明没有找到
	return right // 在right之后插入
}



// 3. 快速排序
func QuickSort(nums []int) {
	quickSortHelper(nums, 0, len(nums)-1)
}

// 3.1 递归版本的快速排序
func quickSortHelper(nums []int, left, right int) {
	//
	if left <= right {
		pivot := getPartition(nums, left, right)
		// 然后对两边继续快排
		quickSortHelper(nums, left, pivot-1)
		quickSortHelper(nums, pivot+1, right)
	}
}
func getPartition(nums []int, left, right int) int {
	// 如何选择基准元素
	// 一般选择最左边元素、中间元素、最右侧元素、1/3处元素、随机元素
	pivot := left + (right-left)/2 // 随便选择一个标准的
	// 下面需要根据pivot位置的值划分数组
	i, j := left, right // 闭区间
	pivotElem := nums[pivot]
	for i <= j { // i==j 有意义
		// 找到一个大于pivot的
		for i <= j && nums[i] <= pivotElem {
			i++
		}
		// 找到一个小于pivotElem
		for i <= j && nums[j] > pivotElem {
			j--
		}
		if i > j {
			break
		}
		// 交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// pivotElem归位
	nums[j] = pivotElem
	return pivot
}

// 4. 归并排序
func MergeSort(nums []int) {
	mergeSort1(nums, 0, len(nums)-1)
}

// func mergeSortHelper(nums []int, left,right int) {
// 	// 对【left,right】归并排序
// 	// 非迭代版本
// 	curSize := 2  //当前排序的尺度  两两排序
// 	n := len(nums)
// 	for curSize <= n {
// 		for i:=0;i<n-curSize;{
// 			mid := i + curSize/2
// 			merge(nums[i: i+mid], nums[i+mid:i+c])
// 		}

// 	}
// }
func mergeSort1(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSort1(nums, left, mid) // 闭区间
	mergeSort1(nums, mid+1, right)
	merge(nums, left, mid, right)
}
func merge(arr []int, start, mid, end int) {
	var tmparr = []int{}
	var s1, s2 = start, mid + 1
	for s1 <= mid && s2 <= end {
		if arr[s1] > arr[s2] {
			tmparr = append(tmparr, arr[s2])
			s2++
		} else {
			tmparr = append(tmparr, arr[s1])
			s1++
		}
	}
	if s1 <= mid {
		tmparr = append(tmparr, arr[s1:mid+1]...)
	}
	if s2 <= end {
		tmparr = append(tmparr, arr[s2:end+1]...)
	}
	for pos, item := range tmparr {
		arr[start+pos] = item
	}
}

// 5. 堆排序
func HeapSort(nums []int) {
	// 声明堆
	h := &IHeap{}
	heap.Init(h)
	// 如堆
	for _, v := range nums {
		heap.Push(h, v)
	}
	// 出队
	for i := 0; i < len(nums); i++ {
		nums[i] = heap.Pop(h).(int)
	}
}

// 构建大根堆、小根堆  元素倒入，再倒出就可以得到有序序列
// 直接调用接口container/heap
type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 6. 选择排序
func SelectSort(nums []int) {
	// 选择排序，每次从未排序部分选择下一个有序的值
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		minIdx := i
		for j := i; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				minIdx = j
			}
		}
		// 将min插入当前位置，也就是交换位置
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

// 桶排序 基数排序 希尔排序
/*
桶内排序
*/
func sortInBucket(bucket []int) { //此处实现插入排序方式，其实可以用任意其他排序方式
	length := len(bucket)
	if length == 1 {
		return
	}

	for i := 1; i < length; i++ {
		backup := bucket[i]
		j := i - 1
		//将选出的被排数比较后插入左边有序区
		for j >= 0 && backup < bucket[j] { //注意j >= 0必须在前边，否则会数组越界
			bucket[j+1] = bucket[j] //移动有序数组
			j--                     //反向移动下标
		}
		bucket[j+1] = backup //插队插入移动后的空位
	}
}

/*
获取数组最大值
*/
func getMaxInArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

/*
桶排序
*/
func BucketSort(arr []int) []int {
	//桶数
	num := len(arr)
	//k（数组最大值）
	max := getMaxInArr(arr)
	//二维切片
	buckets := make([][]int, num)

	//分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num - 1) / max //分配桶index = value * (n-1) /k

		buckets[index] = append(buckets[index], arr[i])
	}
	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sortInBucket(buckets[i])

			copy(arr[tmpPos:], buckets[i])

			tmpPos += bucketLen
		}
	}

	return arr
}

// 基数排序
// 基数排序
// 从个位开始排序，从低到高进行递推
// 高位相同，则顺序不变

// 取得数组中的最大数，并取得位数；
// arr为原始数组，从最低位开始取每个位组成radix数组；
// 对radix进行计数排序（利用计数排序适用于小范围数的特点）；
func Radixsort(arr []int) []int {
	maxValueLen := 0 // 需要循环多少次，以最大数字为准
	for i := 0; i < len(arr); i++ {
		n := len(strconv.Itoa(arr[i])) // 方便起见，数字转字符，再取长度
		if n > maxValueLen {
			maxValueLen = n
		}
	}
	for loc := 1; loc <= maxValueLen; loc++ {
		sort(arr, loc)
	}
	return arr
}

func sort(arr []int, loc int) {
	bucket := make([][]int, 10) // 0~9 总共10个队列

	for i := 0; i < len(arr); i++ {
		ji := digit(arr[i], loc) // 获取对应位的数字
		if bucket[ji] == nil {
			bucket[ji] = make([]int, 0) // 如果 bucket 需要再去初始化
		}

		bucket[ji] = append(bucket[ji], arr[i]) // 将数字 push 进去
	}

	idx := 0
	for i := 0; i <= 9; i++ {
		for j := 0; j < len(bucket[i]); j++ {
			// 遍历二维数组
			arr[idx] = bucket[i][j] //将数字取出来 给原数组重新赋值
			idx++
		}
	}
}

// 数字，右数第几位，从1开始
func digit(num int, loc int) int {
	return num % int(math.Pow10(loc)) / int(math.Pow10(loc-1))
}
