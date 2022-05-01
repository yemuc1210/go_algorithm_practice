package tencent50
/* 简单
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
*/

/*
栈
*/
func isValid(s string) bool {
	n := len(s)
	if n==0 || s==""{
		return true
	}
	//左括号入栈 右括号出
	stack := make([]rune, n)
	stack = append(stack, rune(s[0]))
	//一直字符串只包含括号
	if stack[0]==')' || stack[0]=='}' || stack[0]==']' {
		//右括号
		return false
	}
	//否则第一个是左括号
	stack = stack[:0]  //清空
	for _,ch := range s {
		if ch=='(' || ch=='{' || ch=='[' {
			stack = append(stack, ch)
		}else if (ch==']' && len(stack)>0 && stack[len(stack)-1]=='[') ||
			(ch==')' && len(stack)>0 && stack[len(stack)-1]=='(') || 
			(ch=='}' && len(stack)>0 && stack[len(stack)-1]=='{') {
				//与栈顶符号匹配
				stack = stack[:len(stack)-1]  //弹出栈顶左括号
		}else{ //未遇到左括号  右括号 出错
			return false
		}
	}
	return len(stack) == 0  //栈空则完全匹配
}