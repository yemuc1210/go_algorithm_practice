package lt150

import "strconv"

//输入逆波兰式求，后缀求中缀，栈  专项版36
func evalRPN(tokens []string)int{
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
