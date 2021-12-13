package NC49

/**
  * 最长的括号子串 dp lt32
  * @param s string字符串 
  * @return int整型
*/
/*
题目描述：
给出一个长度为 n 的，仅包含字符 '(' 和 ')' 的字符串，计算最长的格式正确的括号子串的长度。
maxDP
*/
/*
最长  
括号匹配用栈  动态规划
思路1：记录连续出栈的最大次数
	如果连续出栈n次，则最长有效括号为2*n
思路2；先用栈模拟，将无法匹配的括号位置置为1  然后求最长的连续0的长度
思路3：dp
	dp[i]:s[i]结尾的最长有效括号长度
*/
//思路2
func longestValidParentheses( s string ) int {
	if len(s)==0{
		return 0
	}
	length := len(s)
	dp := make([]int,length)
	for i:=1;i<length;i++{
		//遇到右括号 向前匹配
		if s[i]==')'{
			pre := i - dp[i-1] - 1
			//如果是左括号 更新匹配长度
			if pre>=0 && s[pre]=='('{
				dp[i] = dp[i-1] + 2
				//处理独立括号对 （）（）   （）（（））
				if pre > 0{
					dp[i] += dp[pre-1]
				}
			}
		}
	}
	maxDP := dp[0]
	for i:=1;i<len(dp);i++{
		if dp[i]>maxDP {
			maxDP = dp[i]
		}
	}
	return maxDP
}