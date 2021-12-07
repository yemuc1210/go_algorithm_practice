package offerS36

import "strconv"

/*
逆波兰表达式求值  后缀
剑指 Offer II 036. 后缀表达式
根据 逆波兰表示法，求该后缀表达式的计算结果。
有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
示例 1：
	输入：tokens = ["2","1","+","3","*"]
	输出：9
	解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
逆波兰表达式主要有以下两个优点：

    去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
    适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。
*/
func evalRPN(tokens []string) int {
	stack := []int{}

	for _,token := range tokens{
		val,err := strconv.Atoi(token)  //转为整数
		if err == nil {
			stack = append(stack, val)
		}else{
			//是运算符
			//弹出两个
			num1,num2:= stack[len(stack)-2],stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			switch token{
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}