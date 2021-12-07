package NC52
/**
  * 有效的括号序列  栈即可
  * @param s string字符串 
  * @return bool布尔型
*/
func isValid( s string ) bool {
    // write code here
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 && (s[0]==')'||s[0]=='}'||s[0]==']'){
		return false
	}
	stack := []byte{}
	//遇到左括号压栈
	for idx:=0; idx < len(s) ; idx++ {
		if s[idx] == '(' || s[idx] == '{' || s[idx] == '[' {
			stack = append(stack, s[idx])
		}else if (s[idx] == ')' && len(stack)>0 && stack[len(stack)-1] == '(') ||
			(s[idx] == '}' &&  len(stack)>0 && stack[len(stack)-1] == '{') ||
			(s[idx] == ']' && len(stack)>0 && stack[len(stack)-1] == '[') {
			//合法 弹出栈顶
			stack = stack[:len(stack)-1]   //可能非法
		}else {
			//非法字符 跳过
			return false
		}
	}
	return len(stack) == 0
}