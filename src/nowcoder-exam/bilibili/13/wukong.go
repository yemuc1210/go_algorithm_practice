package main

import (
	"fmt"
	"sort"
)

/*
描述
打败魔人布欧以后，孙悟空收了n个徒弟，每个徒弟战斗力各不相同。
他教导所有的徒弟和体术，合体后战斗力为原战斗力相乘。
任何两个徒弟都可以合体，所以一共有n*(n-1)/2种合体徒弟。
有一天，他想考验一下孙悟天战斗力如何，
希望在所有n*(n-1)/2种合体徒弟中选择战斗力第k高的，与孙悟天对战。
可是孙悟空徒弟太多了，他已然懵逼，于是找到了你，请你帮他找到对的人。

输入描述：
第一行两个int。徒弟数量：n <= 1*10^6；战斗力排名:k <= n*(n-1)/2
第二行空格分隔n个int，表示每个徒弟的战斗力。
输出描述：
战斗力排名k的合体徒弟战斗力。
*/

func main() {
	var n,k int
	fmt.Scanf("%d %d",&n,&k)

	nums := make([]int64,n)
	for i:=0;i<n;i++{
		fmt.Scanf("%d%c",&nums[i])
	}
	// fmt.Println(nums)
	//得到合体的战斗力数据集合
	sets := make(map[int64]int64)
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			sets[nums[i]*nums[j]] ++
		}
	}
	//sets转为数组 然后二分查找
	arr := []int64{}
	for k,v := range sets {
		for v>0 {
			arr = append(arr, k)
			v -- 
		} 
	}
	//问题变为查找第k大
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(findKthLargest2(arr,k))
}


//基于堆排序的方案
func findKthLargest2(nums []int64, k int) int64 {
    heapSize := len(nums)
    buildMaxHeap(nums, heapSize)  //建立大根堆
    for i := len(nums) - 1; i >= len(nums) - k + 1; i-- {  //做k-1次删除操作后，堆顶元素就是
        nums[0], nums[i] = nums[i], nums[0]
        heapSize--
        maxHeapify(nums, 0, heapSize)
    }
    return nums[0]
}

func buildMaxHeap(a []int64, heapSize int) {
    for i := heapSize/2; i >= 0; i-- {
        maxHeapify(a, i, heapSize)
    }
}

func maxHeapify(a []int64, i, heapSize int) {
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