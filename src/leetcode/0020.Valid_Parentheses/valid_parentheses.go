/*
有效括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
*/

package leetcode

func isValid(s string) bool {
	//之前遇到过吧，用栈结构     好像编译原理里面也遇到过，状态机
	n := len(s)    //调用函数开销大，还是存储变量开销大啊，应该不比每次都len(s)吧
	if n==0 ||s=="" {  //两个等价的，其实判断到第一个就出结果了，
		return true
	}
	//遇到左括号入栈，遇到右括号开始出栈   因为就那几种括号，就直接列举了
	
	stack := make([]rune,n)
	stack = append(stack, rune(s[0]))
	//字符串只包含括号，所以第一个必定也是括号了    也可能是右括号。。。
	if stack[0]==')' || stack[0]=='}' || stack[0]==']' {
		return false
	}
	stack = stack[:0]
	//只包含括号，规律就是中间那一组括号是相邻的
	for _,ch := range s {
		if (ch == '[') || (ch == '(') || (ch == '{') {   //左括号入栈
			stack = append(stack, ch)
		}else if ( ch==']' && len(stack)>0 && stack[len(stack)-1]=='[') ||
			( ch==')' && len(stack)>0 && stack[len(stack)-1]=='(') ||
			(ch=='}' && len(stack)>0 && stack[len(stack)-1]=='{'){
				//弹出
				stack = stack[:len(stack)-1]//弹出左括号，右括号也没有入队
		}else{   //并没有遇到右括号，遇到右括号时就应该执行弹出了，并且前一个应该是匹配的
			return false
		}
	}
	return len(stack)==0  //所有的都匹配，栈空  
}