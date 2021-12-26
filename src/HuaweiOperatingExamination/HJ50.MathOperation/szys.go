package main
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
    "unicode"
)
//四则运算 输入表达式 求值  规定表达式一定合法
/*
思路：
1 规定符号优先级
2 遇到数字 查后续符号 或者单设一个符号栈
	当符号比当前符号优先级更高，做出栈操作并计算
3 遇到符号  一般就是括号
	左括号入栈
	右括号连续出栈 计算 直到遇到对应的左括号
*/


func main(){
    var params []string
    input := bufio.NewScanner(os.Stdin)
    for input.Scan(){
        params = append(params, input.Text())
    }
    for loop:= 0; loop< len(params); loop++{
        tmp := strings.Replace(strings.Replace(params[loop], "]", ")", -1), "[", "(", -1)
        tmp = strings.Replace(strings.Replace(tmp, "}", ")", -1), "{", "(", -1)
        res := OutPrint(tmp)
        fmt.Println(res)
    }
}

func OutPrint(s string) int {
    bytes := make([]byte, len(s))
    copy(bytes, s)
    return subCal(&bytes)
}

func subCal(bytes *[]byte) int {
    stack := []int{}
    num := 0
    var sigh byte = '+'
    for len(*bytes) > 0 {
        item := (*bytes)[0]
        *bytes = (*bytes)[1:]
        if unicode.IsDigit(rune(item)) {
            val, _ := strconv.Atoi(string(item))
            num = num*10 + val
        }
        if item == '(' {
            num = subCal(bytes)
        }
        if (!unicode.IsDigit(rune(item)) && item != ' ') || len(*bytes) == 0 {
            switch sigh {
                case '+':
                    stack = append(stack, num)
                case '-':
                    stack = append(stack, -num)
                case '*':
                    num *= stack[len(stack)-1]
                    stack = stack[:len(stack)-1]
                    stack = append(stack, num)
                case '/':
                    num = stack[len(stack)-1]/num
                    stack = stack[:len(stack)-1]
                    stack = append(stack, num)
            }
        sigh = item
        num = 0
        }
        if item == ')' {
            break
        }
    }
    sum := 0
    for i := 0; i < len(stack); i++ {
        sum += stack[i]
    }
    return sum
}

