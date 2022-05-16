package main

import "fmt"

// 分割数组的方案数
// 下标0开始，长度n
func waysToSplitArray(nums []int) int {
	n := len(nums)
	// 合法分割
	// 前i+1元素和大于剩余的n-i-1元素和
	totalSum := 0 
	// 前缀和方法
	// prefixSum[i] 前i+1个元素的和
	prefixSum := make([]int,n)
	for i,num := range nums {
		totalSum += num
		prefixSum[i] = totalSum
	}
	// 下标右边至少一个元素
	var res int
	for i:=0;i<n-1;i++ {
		if prefixSum[i] >= totalSum-prefixSum[i] {
			res ++
		}
	}
	return res
}

func main() {
	// nums := []int{10,4,-8,7}
	nums := []int{2,3,1,0}
	res := waysToSplitArray(nums)
	fmt.Println(res)
}
