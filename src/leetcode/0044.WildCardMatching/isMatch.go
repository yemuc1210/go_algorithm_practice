package lt44

/*  困难  lt10 正则表达式匹配
44. 通配符匹配
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
	'?' 可以匹配任何单个字符。
	'*' 可以匹配任意字符串（包括空字符串）。
	两个字符串完全匹配才算匹配成功。
说明:
	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
*/
/*
	dp:
	dp[i][j] 表示s-i p-j是否匹配
	初始化：
		dp[0][0] true
		第一行dp[0][j]  s 为空 p若开始为*则true
		第一列dp[i][0]   falase
	状态转移方程：
		(s[i]==p[j] || p[j]=='?' )  && dp[i-1][j-1]
			dp[i][j] = true
		p[j] = '*' && (dp[i-1][j]=true || dp[i][j-1]=true)
			dp[i][j] = true
				dp[i-1][j] 表示*代表的是空字符  ab ab*
				dp[i][j-1] 表示*代表非空  abcd  ab*
*/
func isMatch(s string, p string) bool {
	sn,pn := len(s),len(p)
	dp := make([][]bool,sn+1)
	for i:=range dp{
		dp[i] = make([]bool, pn+1)
		if i==0 {
			dp[i][0] = true
		}
	}
	//初始化第一行
	for j:=1;j<=pn;j++{
		//若p开始为*则true
		if p[j-1]=='*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	//
	for i:=1;i<=sn;i++{
		for j:=1;j<=pn;j++{
			if s[i-1]==p[j-1] || p[j-1]=='?'{
				dp[i][j] = dp[i-1][j-1]
			}else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[sn][pn]
}