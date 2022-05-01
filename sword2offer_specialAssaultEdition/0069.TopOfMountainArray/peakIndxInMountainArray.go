package offerS69

import "fmt"

/*
剑指 Offer II 069. 山峰数组的顶部
符合下列属性的数组 arr 称为 山峰数组（山脉数组） ：

arr.length >= 3
存在 i（0 < i < arr.length - 1）使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
给定由整数组成的山峰数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i ，即山峰顶部。

山峰数组中
*/
func peakIndexInMountainArray(arr []int) int {
	//山峰左右大小关系不一致
	//已知长度大于等于3
	if arr[0] > arr[1] {
		return 0
	}
	for i:=1;i<=len(arr)-2;i++ {
		fmt.Println(arr[i-1],arr[i],arr[i+1])
		fmt.Printf("%v   %v\n",arr[i]>arr[i-1],arr[i]<arr[i+1])
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return len(arr)-1
}