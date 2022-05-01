package lt224

import "container/list"

/*
224. 基本计算器
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。



示例 1：

输入：s = "1 + 1"
输出：2
示例 2：

输入：s = " 2-1 + 2 "
输出：3
示例 3：

输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23
*/
/*
计算表达式的，一般使用 逆波兰式
首先 将中缀表达式转为后缀表达式 直接中缀计算
*/
func calculate(s string) int{
	stack := list.New()
	sign := 1// 正负号
	idx := 0 //当前需要处理的字符下标
	res := 0
	for idx < len(s) {
		//可能为空格
		if s[idx] == ' '{
			idx ++
		}else if s[idx]>='0' && s[idx]<='9' {
			//获得一段数字
			base,val := 10, int(s[idx]-'0')
			for idx+1<len(s) && s[idx+1]>='0'&&s[idx+1]<='9' {
				val = val*base + int(s[idx+1] -'0')
				idx ++
			}
			res += val*sign
			idx ++ 
		}else if s[idx] =='+' {
			sign = 1
			idx ++
		}else if s[idx]=='-'{
			sign = -1
			idx ++
		}else if s[idx] == '(' {
			//之前的结果压栈
			stack.PushBack(res)
			stack.PushBack(sign)
			//开始新的计算
			res = 0
			sign = 1
			idx ++
		}else if s[idx] == ')' {
			// 新的计算结果 * 前一个加减状态 + 之前计算结果
			res = res*stack.Remove(stack.Back()).(int) + stack.Remove(stack.Back()).(int)
			idx++
		}
	}
	return res
}