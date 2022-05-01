package lt516
/*
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列


dp[i][j]  表示s[i...j]是最长的回文
dp[i][i] = 1
dp[i][j] = dp[i+1][j-1] + 2   s[i]=s[j]
dp[i][j] = max(dp[i+1][j],  dp[i][j-1])  s[i]!=s[j]

*/
func longestPalindromeSubseq(s string) int {
	lenS := len(s)
	dp := make([][]int,lenS)
	for i:=0;i<lenS;i++{
		dp[i] = make([]int, lenS)
	}

	for i:=0;i<lenS;i++{
		dp[i][i] = 1   //单个字符是回文
	}
	for i:=lenS-1;i>=0;i--{
		for j:=i+1;j<lenS;j++{
			if s[i]==s[j]{
				dp[i][j] = dp[i+1][j-1] + 2   //长度+2
			}else{
				dp[i][j] = max(dp[i+1][j],  dp[i][j-1])
			}
		}
	}
	return dp[0][lenS-1]
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}