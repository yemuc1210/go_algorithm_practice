package lt60

import "strconv"

/* 困难
60. 排列序列
给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

示例 1：
输入：n = 3, k = 3
输出："213"
示例 2：
输入：n = 4, k = 9
输出："2314"
示例 3：
输入：n = 3, k = 1
输出："123"
提示：
1 <= n <= 9
1 <= k <= n!
*/
/*
排列 第K个  数字[1,2,3,...,n]
dfs
*/
func getPermutation(n,k int) string{
	if k==0 {
		return ""
	}
	used,p,res := make([]bool, n), []int{}, ""
	findPermutation(n,0,&k,p,&res,&used)
	return res
}

func findPermutation(n,index int, k *int, p []int, res *string, used *[]bool) {
	if index==n {
		*k--
		if *k==0{
			for _,v := range p {
				*res += strconv.Itoa(v+1)
			}
		}
		return 
	}
	for i:=0;i<n;i++{
		if !(*used)[i] {
			(*used)[i]=true
			p = append(p, i)
			findPermutation(n,index+1,k,p,res,used)
			p = p[:len(p)-1] //回溯
			(*used)[i]=false
		}
	}
}