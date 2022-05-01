package lt97

/* lt97   dp
剑指 Offer II 096. 字符串交织
给定三个字符串 s1、s2、s3，请判断 s3 能不能由 s1 和 s2 交织（交错） 组成。
两个字符串 s 和 t 交织 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
	s = s1 + s2 + ... + sn
	t = t1 + t2 + ... + tm
	|n - m| <= 1
	交织 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
提示：a + b 意味着字符串 a 和 b 连接。

新串中子串的相对顺序不变
*/
/*
寻找公共字串  寻找一次 然后在新串中移除这个子串   不行，匹配的字符顺序也很重要
	本来都可以匹配的，结果第一个某个字符匹配错位置

dp[i][j]=(dp[i-1][j] and s1[i-1]==s3[i+j-1])or(dp[i][j-1] and s2[j-1]==s3[i+j-1])
dp[i][j]状态代表当前s1[:i] s2[:j] 能否交叉组成 s3[:i+j-1]
*/
func isInterleave(s1 string, s2 string, s3 string) bool {
	m,n,k := len(s1),len(s2),len(s3)
	if m+n != k {
		return false  //长度不匹配
	}

	dp := make([][]bool,m+1)
	for i:= range dp{
		dp[i] = make([]bool, n+1)
	}

	//dp
	dp[0][0] = true  //都为nil
	for i:=1;i<m+1;i++{  //初始化
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]  //s2空
	}
	for i:=1;i<n+1;i++{  //初始化 如果s1为空的情况下,初始化dp[0][j]
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]  //s2空
	}
	//一般情况
	for i:=1;i<m+1;i++{
		for j:=1;j<n+1;j++{
			//要么s1[i-1]==s3[i+j-1]且之前的都匹配上了要么s2[j-1]==s3[i+j-1]之前的也都匹配上了
			dp[i][j]=(dp[i-1][j] &&s1[i-1]==s3[i+j-1]) || (dp[i][j-1] && s2[j-1]==s3[i+j-1])
		}
	}
	return dp[m][n]
}
