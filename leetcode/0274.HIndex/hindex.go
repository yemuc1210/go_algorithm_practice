package lt274

import "sort"

/*
给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。
h 指数的定义：h 代表“高引用次数”（high citations），
一名科研人员的 h 指数是指他（她）的 （N 篇论文中）总共有 h 篇论文分别被引用了至少 h 次。
且其余的 N - h 篇论文每篇被引用次数 不超过 h 次。

例如：某人的 h 指数是 20，这表示他已发表的论文中，每篇被引用了至少 20 次的论文总共有 20 篇。。
*/
func hIndex(citations []int) int {
	sort.Ints(citations)
	hindex := 0 
	for i:=len(citations)-1 ;i>=0;i--{
		if citations[i]>=len(citations)-i{
			hindex++
		}else{
			break
		}
	}
	return hindex

}

func hIndex1(citations []int)int{
	quickSort(citations,0,len(citations)-1)
	hindex := 0 
	for i:=len(citations)-1 ;i>=0;i--{
		if citations[i]>=len(citations)-i{
			hindex++
		}else{
			break
		}
	}
	return hindex
}
func quickSort(a []int, low,high int){
	if low>=high{
		return
	}
	p := partition(a,low,high)
	quickSort(a,low,p-1)
	quickSort(a,p+1,high)
}
func partition(a []int, low,high int)int{
	pivot := a[high]
	i:=low-1
	for j:=low;j<high;j++{
		if a[j]<pivot{
			i++
			a[j],a[i] = a[i],a[j]
		}
	}
	a[i+1],a[high] = a[high],a[i+1]
	return i+1

}
// 解法一
func hIndex2(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}
	count := 0
	for i := n; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}
	return 0
}