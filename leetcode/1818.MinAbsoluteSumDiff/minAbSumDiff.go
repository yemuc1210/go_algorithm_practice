package lt1818

import (
	// "fmt"
	// "math"
	"sort"
)

/*
给你两个正整数数组 nums1 和 nums2 ，数组的长度都是 n 。

数组 nums1 和 nums2 的 绝对差值和 定义为所有 |nums1[i] - nums2[i]|（0 <= i < n）的 总和（下标从 0 开始）。

你可以选用 nums1 中的 任意一个 元素来替换 nums1 中的 至多 一个元素，以 最小化 绝对差值和。

在替换数组 nums1 中最多一个元素 之后 ，返回最小绝对差值和。因为答案可能很大，所以需要对 109 + 7 取余 后返回。

|x| 定义为：
如果 x >= 0 ，值为 x ，或者
如果 x <= 0 ，值为 -x
*/
// type pairs struct{
// 	index int
// 	diff int
// }
// func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
// 	diffs := []pairs{}
// 	for i:=0;i<len(nums1);i++{
// 		diffs = append(diffs, pairs{index: i,diff:abs(nums1[i],nums2[i])})
// 	}
// 	//根据diff值排序
// 	merge_sort(diffs,0,len(diffs)-1)
// 	maxPair := diffs[len(diffs)-1]
// 	fmt.Println(diffs)
// 	if maxPair.diff == 0{
// 		return 0
// 	}
// 	minA := maxPair.diff
// 	for i:=0;i<len(nums1);i++{
// 		if i == maxPair.index{
// 			continue
// 		}
// 		//从nums1中找到abs(b-nums1)最小的
// 		// fmt.Printf("nums2[maxIndex]=%v,nums1[i]=%v,i=%v,minA=%v\n",nums2[maxPair.index],nums1[i],i,minA)
// 		if minA > abs(nums2[maxPair.index],nums1[i]){
// 			minA = abs(nums2[maxPair.index], nums1[i])
// 		}
// 	}

// 	diffs[len(diffs)-1].diff = minA  //可能存在未修改的情况
// 	//求和
// 	sum := 0
// 	for i:=0;i<len(diffs);i++{
// 		sum = (sum + diffs[i].diff)
// 		fmt.Printf("sum=%v\n",sum)

// 	}
// 	return sum%(int(math.Pow(10,9)+7))
// }

// func merge_sort(A []pairs, start, end int) int {
//     if start >= end {
//         return 0
//     }
//     mid := start + (end - start)>>1
//     left := merge_sort(A, start, mid)
//     right := merge_sort(A, mid+1, end)
//     cross := merge(A, start, mid, end)
//     return left + right + cross
// }
// func merge (A []pairs, start, mid, end int) int{
//     Arr := make([]pairs, end-start+1)
//     p, q, k, count := start, mid+1, 0, 0
//     for i := start; i <= end; i++ {
//         if p > mid {
//             Arr[k] = A[q]
//             q++
//         } else if q > end {
//             Arr[k] = A[p]
//             p++
//         } else if A[p].diff <= A[q].diff {
//             Arr[k] = A[p]
//             p++
//         } else {
//             count += mid - p + 1
//             Arr[k] = A[q]
//             q++
//          }
//         k++
//     }
//     copy(A[start:end+1], Arr)
//     return count
// }

// func abs(a,b int)int{
// 	if a-b>=0{
// 		return a-b
// 	}
// 	return b-a
// }
//最大化|nums1[i]-nums2[i]| - |nums1[j]-nums2[i]|
func minAbsoluteSumDiff(nums1, nums2 []int) int {
	rec := append(sort.IntSlice(nil), nums1...)
	rec.Sort()
	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)
		sum += diff
		j := rec.Search(v)
		if j < n {
			maxn = max(maxn, diff-(rec[j]-v))
		}
		if j > 0 {
			maxn = max(maxn, diff-(v-rec[j-1]))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
