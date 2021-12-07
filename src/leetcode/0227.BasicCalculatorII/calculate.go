package lt227
/* lt224
227. 基本计算器 II
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
整数除法仅保留整数部分。
*/
/*
比224多出乘除
先计算乘除 然后问题变成只有加减的
*/
func calculate(s string) (ans int) {
    stack := []int{}
    preSign := '+'
    num := 0
    for i, ch := range s {
        isDigit := '0' <= ch && ch <= '9'
        if isDigit {
            num = num*10 + int(ch-'0')
        }
        if !isDigit && ch != ' ' || i == len(s)-1 {
            switch preSign {
            case '+':
                stack = append(stack, num)
            case '-':
                stack = append(stack, -num)
            case '*': //遇到乘除的 与栈顶相乘除  再放入
                stack[len(stack)-1] *= num
            default:
                stack[len(stack)-1] /= num
            }
            preSign = ch
            num = 0
        }
    }
    for _, v := range stack {
        ans += v
    }
    return
}
