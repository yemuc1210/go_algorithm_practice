package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//简单表达式计算

func main(){
	//输入有多行字符串  每行是一个表达式
	reader := bufio.NewScanner(os.Stdin)
	for  reader.Scan(){
		expr := reader.Text()
		expr = strings.Trim(expr,"\n")
		expr = strings.TrimSpace(expr)
		// fmt.Printf("%s, %#v,  %T\n",expr,expr,expr)
		if expr == "END" {
			break
		}
		//表达式字符串：只包含非负整数 加法减法乘法  没有括号
		fmt.Println(calculate(expr))
	}
}

func calculate(s string) int {
	//因为只包含非负整数+没有除法 所有结果也是整数
	//计算方法是遇到加法压入栈，遇到负号将相反数压入栈，（第一个前面补个+）
	//遇到乘法出栈计算，结果压入栈
	s = "+" + s
	stack := []int{}
	idx := 0 
	for idx<len(s) {
		if s[idx] == '+' {
			//入栈
			// num,err := strconv.Atoi(s[idx+1:idx+2])  //错误 数字不是只有一位得
			num := 0
			for idx=idx+1;idx<len(s) && (s[idx]!='+' && s[idx]!='-' && s[idx]!='*'); idx++{
				cur,err := strconv.Atoi(s[idx:idx+1])
				if err!=nil {
					panic(err)
				}
				num = num*10 + cur 
			}  //遇到符号位停止  得到完整得数字
			// fmt.Println(num)
			stack = append(stack, num)
		}else if s[idx] == '-' {
			//负数入栈
			num := 0
			for idx=idx+1;idx<len(s) && (s[idx]!='+' && s[idx]!='-' && s[idx]!='*'); idx++{
				cur,err := strconv.Atoi(s[idx:idx+1])
				if err!=nil {
					panic(err)
				}
				num = num*10 + cur 
			}  //遇到符号位停止  得到完整得数字
			// num,err := strconv.Atoi(s[idx+1:idx+2])
			
			// fmt.Println(0-num)
			stack = append(stack, 0-num)
		}else if s[idx] == '*' {
			//结果入栈
			num := 0
			for idx=idx+1;idx<len(s) && (s[idx]!='+' && s[idx]!='-' && s[idx]!='*'); idx++{
				cur,err := strconv.Atoi(s[idx:idx+1])
				if err!=nil {
					panic(err)
				}
				num = num*10 + cur 
			}  //遇到符号位停止  得到完整得数字
			lastNum := stack[len(stack)-1]
			stack = stack[:len(stack)-1]  //出栈
			// fmt.Println(num,lastNum,num*lastNum)
			stack = append(stack, num*lastNum)  //结果再入栈
		}
		// idx += 2  //跳过符号位和数字  第一个数补0省去第一个数得判断
	}
	//将剩余数字相加
	res := 0
	for len(stack) > 0 {
		res += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}