package main

import (
	"fmt"
	"sort"
)


func numFactoredBinaryTrees(arr []int) int {
	// dp
	// 最大值一定为树根
	// dp[v]： 以v为根的树种类数
	// 如果有树根节点有孩子x y 满足x*y=v 则有dp[x]*dp[y]种方式
	n:=len(arr)
	dp := make([]int, n)
	// 不妨先对arr排序
	sort.Slice(arr,func(i, j int) bool {
		return arr[i]<arr[j]
	})
	index := make(map[int]int)
	for i,v := range arr{
		index[v] = i
		dp[i] = 1
	}
	// 
	Mod := 1000000007
	for i:=0;i<n;i++{
		for j:=0;j<i;j++{
			if arr[i]%arr[j] == 0 {
				right := arr[i]/arr[j]   // arr[j] right为左右孩子
				if _,ok := index[right];ok {
					// 存在，可以构建
					dp[i] = (dp[i]+dp[j]*dp[index[right]])%Mod
				}
			}
		}
	}
	res := 0
	for _,v := range dp{
		res += v
	}
	return res%Mod
}

func main() {
	arr := []int{2,4,5,10}
	res := numFactoredBinaryTrees(arr)
	fmt.Println(res)
}