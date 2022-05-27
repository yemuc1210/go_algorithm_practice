package main

// 栈从空到下一次空  是一个原语。且应该舍去首尾
func removeOuterParentheses(s string) string {
    var ans, st []rune
    for _, c := range s {
        if c == ')' {
            st = st[:len(st)-1]
        }
        if len(st) > 0 {
            ans = append(ans, c)
        }
        if c == '(' {
            st = append(st, c)
        }
    }
    return string(ans)
}
