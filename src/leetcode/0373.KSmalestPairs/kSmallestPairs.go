package lt373


/*
剑指 Offer II 061. 和最小的 k 个数对
给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
*/
import (
    "container/heap"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    arrays := generateTwoArrs(nums1, nums2)
    if len(arrays) <= k {
        return arrays
    }

    heap.Init(&arrays)
    result := make(DoubleArray, k)
    for i := 0; i < k; i++ {
        result[i] = (heap.Pop(&arrays)).([]int)
    }

    return result
}

func generateTwoArrs(nums1, nums2 []int) DoubleArray {
    arrs := make(DoubleArray, 0, len(nums1) * len(nums2))
    for n := 0; n < len(nums1); n++ {
        for m := 0; m < len(nums2); m++ {
            arrs = append(arrs, []int{nums1[n], nums2[m]})
        }
    }

    return arrs
}

// 实现container/heap的接口
type DoubleArray [][]int

func (doubleArray *DoubleArray) Push(x interface{}) {
    *doubleArray = append(*doubleArray, x.([]int))
}

func (doubleArray *DoubleArray) Pop() interface{} {
    length := len(*doubleArray)
    val := (*doubleArray)[length - 1]
    *doubleArray = (*doubleArray)[: length - 1]

    return val
}

func (doubleArray DoubleArray) Len() int {
    return len(doubleArray)
}

func (doubleArray DoubleArray) Less(i int, j int) bool {
    sum1 := doubleArray[i][0] + doubleArray[i][1]
    sum2 := doubleArray[j][0] + doubleArray[j][1]
    return sum1 <= sum2
}

func (doubleArray DoubleArray) Swap(i int, j int) {
    doubleArray[i], doubleArray[j] = doubleArray[j], doubleArray[i]
}


