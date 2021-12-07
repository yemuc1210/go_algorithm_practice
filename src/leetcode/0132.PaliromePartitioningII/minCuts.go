package lt132

import "fmt"

/*lt132  dp 腾讯2020-08-23笔试第5题
剑指 Offer II 094. 最少回文分割
给定一个字符串 s，请将 s 分割成一些子串，使每个子串都是回文串。

返回符合要求的 最少分割次数 。
*/
/*
以某个字符为中心的最长回文串
*/
func minCut(s string) int {
	if  len(s) <= 1{
		return 0
	}
	length := len(s)
	dp :=make([]int,length)
	//初始化
	for i:=range dp{
		dp[i] = length-1  //最大分割次数就是长度减一
	}
	for i:=0;i<length;i++{

		//奇数回文串
		helper(s,i,i,dp) // 以当前字符为中心扩展
		fmt.Println(dp)
		//偶数
		helper(s,i,i+1,dp)
		// fmt.Println(dp)
	}
	return dp[length-1]
}

func helper(s string, i,j int, dp []int){
	length := len(s)
	for i>=0 && j< length && s[i] == s[j] {
		//回文
		if i == 0 {
			dp[j] = min(dp[j],  0)
		}else{
			dp[j]  = min(dp[j], dp[i-1]+1)
		}
		i--  //左指针  向左扩展
		j++  //右指针
	}
	fmt.Println(dp)
}

func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}
