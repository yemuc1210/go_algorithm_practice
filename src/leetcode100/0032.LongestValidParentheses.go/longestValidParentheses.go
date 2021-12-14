package lt32
/* 困难
32. 最长有效括号
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
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
func longestValidParentheses(s string)int{
	stack := []int{}
	mark := make([]int, len(s))
	length,res := 0,0

	for i:=0;i<len(s);i++{
		if s[i]=='(' {
			stack = append(stack, i)
		}else{
			if len(stack)==0{//多余的右括号
				mark[i]=1
			}else{
				stack = stack[:len(stack)-1] //出战
			}
		}
	}
	//未匹配的左括号标记
	for len(stack)>0{
		mark[stack[len(stack)-1]]=1
		stack = stack[:len(stack)-1]
	}
	//寻找最大连续0
	for i:=0;i<len(mark);i++{
		if mark[i]==1 {
			length = 0
			continue
		}
		length ++
		res = max(res,length)
	}
	return res
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}


//dp  dp[i]:s[i]结尾的最长有效括号长度
func longestValidParentheses1(s string)int{
	if len(s)==0{
		return 0
	}
	length := len(s)
	dp := make([]int,length)
	for i:=1;i<length;i++{
		//遇到右括号 向前匹配
		if s[i]==')'{
			pre := i - dp[i-1] - 1   //dp[i-1]以s[i-1]结尾最长有效括号 所以 s[i]继续往前找
			//如果是左括号 更新匹配长度
			if pre>=0 && s[pre]=='('{
				dp[i] = dp[i-1] + 2
				//处理独立括号对 （）（）   （）（（））
				if pre > 0{  //重点
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
