package NC44

/**
 * //通配符匹配 lt44
 * ?单个字符  *任意字符
 * @param s string字符串 
 * @param p string字符串 
 * @return bool布尔型
*/
func isMatch( s string ,  p string ) bool {
	sn,pn := len(s),len(p)
	dp := make([][]bool,sn+1)
	for i:=range dp{
		dp[i] = make([]bool, pn+1)
		if i==0 {
			dp[i][0] = true
		}else {
			dp[i][0] = false   //第一列
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