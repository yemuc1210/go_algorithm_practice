package lt786

import "sort"

/*困难
786. 第 K 个最小的素数分数
给你一个按递增顺序排序的数组 arr 和一个整数 k 。
数组 arr 由 1 和若干 素数  组成，且其中所有整数互不相同。

对于每对满足 0 < i < j < arr.length 的 i 和 j ，可以得到分数 arr[i] / arr[j] 。
那么第 k 个最小的分数是多少呢?  以长度为 2 的整数数组返回你的答案,
	这里 answer[0] == arr[i] 且 answer[1] == arr[j] 。

示例 1：
	输入：arr = [1,2,3,5], k = 3
	输出：[2,5]
	解释：已构造好的分数,排序后如下所示:
		1/5, 1/3, 2/5, 1/2, 3/5, 2/3
		很明显第三个最小的分数是 2/5
*/
/*
优先队列题不会做啊  得练  还有并查集等
自定义排序，将n(n-1)/2个分数排序
分数比较方式：a/b < c/d    ： a*d < b*c
*/
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)	
	type pair struct { x,y int}
	frac := make([]pair,0, n*(n-1)/2)
	for i,x := range arr {
		for _,y := range arr[i+1:]{
			frac = append(frac, pair{x,y})
		}
	}
	sort.Slice(frac,func(i, j int) bool {
		a,b := frac[i],frac[j]
		return a.x * b.y < a.y * b.x
	})
	return []int{frac[k-1].x, frac[k-1].y}
}