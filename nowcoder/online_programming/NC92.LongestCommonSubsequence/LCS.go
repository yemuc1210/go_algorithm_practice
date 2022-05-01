package NC92

/**
 * longest common subsequence dp  输出最长公共子序列
 * @param s1 string字符串 the string
 * @param s2 string字符串 the string
 * @return string字符串
*/
func LCS( s1 string ,  s2 string ) string {
    // write code here 1 2 3 4 表示各个方向的贡献
	m,n := len(s1),len(s2)
	dp := make([][]int,m+1)
	for i:=range dp{
		dp[i] = make([]int, n+1)
	}
	//dp
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			if s1[i-1] == s2[j-1]{
				dp[i][j] = dp[i-1][j-1] + 1
			}else{
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	//dp[m][n] 
	if dp[m][n] == 0 { // 若没有最长公共子序列
        return "-1"
    }
	var res string
	for i,j :=m,n; dp[i][j]>0; {
		if s1[i-1] == s2[j-1] { //反推公式中相等的场景
            // 该值一定是被选取到的，根据状态转移公式，当时两条字符串的下标都前进一位
            res = string(s1[i-1]) + res // 合并时注意添加进res前面
            i , j = i-1 , j-1 // 反推公式，两个下标都后退一位
        } else if dp[i-1][j] >= dp[i][j-1] { // 当时dp[i][j]结果来自dp[i-1][j]
            i = i-1          // 反推公式，行下标即s1下标i后退一位
        } else if dp[i-1][j] < dp[i][j-1] { // 当时dp[i][j]结果来自dp[i][j-1]
            j = j-1          // 反推公式，列下标即s2下标j后退一位
        }
	}
    return res
}
func max(a,b int) int {
	if a > b {
	  return a
	}
	return b
  }