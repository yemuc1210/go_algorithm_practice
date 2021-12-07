package tencent50
/*中等
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。
*/
/*
1 dp   以当前字符为中心扩散
2 滑动窗口
3 dp
*/
func longestPalindrome(s string) string {
	res,dp := "", make([][]bool, len(s))
	//初始化dp
	for i:= range dp{
		dp[i] = make([]bool, len(s))
	}

	//dp[i][j]表示 s[i:j]是否为回文
	n := len(s)
	for i:=n-1;i>=0;i-- {  //从后往前  判断s[i:n-1]是否是回文
		for j:=i;j<n;j++{  
			dp[i][j] = (s[i] == s[j]) && ((j-i)<3 || dp[i+1][j-1])
			if dp[i][j] && (res=="" || j-i+1 > len(res)) {
				//更新更大的回文串
				res = s[i:j+1]
			}
		}
	}
	return res
}