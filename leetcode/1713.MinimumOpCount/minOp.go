package lt1713

import "sort"

/*
得到子序列的最少操作次数
target 长度为n,arr长度为m，最少添加元素为n-公共子序列长度

求最长公共子序列

已知tatget元素互不相同，所以可以用哈希表记录元素下标，将arr元素映射到下标上
对于不存在与target的元素，其必然不在公共子序列中，将其忽略


*/
func minOperations(target, arr []int) int{
	n := len(target)
	pos := make(map[int]int,n)
	for i,val := range target {
		pos[val] = i
	}	
	d := []int{}
	for _,val := range arr{
		if idx,has := pos[val];has {
			if p := sort.SearchInts(d,idx);p<len(d){
				d[p] = idx
			}else{
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}