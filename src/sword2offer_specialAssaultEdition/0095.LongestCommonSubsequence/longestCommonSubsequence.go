package offerS95
/* lt1143   dp
剑指 Offer II 095. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	m,n := len(text1),len(text2)
	dp := make([][]int,m+1)
	for i:= range dp{
		dp[i] = make([]int, n+1)
	}

	//dp dp[i][j]表示 t1[0:i]与t2[0:j]的最长公共子序列长
	for i,c1 := range text1{
		for j,c2 := range text2{
			if c1==c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			}else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}